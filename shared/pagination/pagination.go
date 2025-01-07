package pagination

import (
	"fmt"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

var (
	// DefaultPageSize specifies the default page size
	DefaultPageSize = 20
	// MaxPageSize specifies the maximum page size
	MaxPageSize = 100
	// PageVar specifies the query parameter name for page number
	PageVar = "page"
	// PageSizeVar specifies the query parameter name for page size
	PageSizeVar = "per_page"
)

// Pages represents a paginated list of data items.
type Pages struct {
	Pagination
}

// Pagination contains pagination information
type Pagination struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	PageCount  int `json:"total_page"`
	TotalCount int `json:"total_data"`
}

// New creates a new Pages instance.
// The page parameter is 1-based and refers to the current page index/number.
// The perPage parameter refers to the number of items on each page.
// And the total parameter specifies the total number of data items.
// If total is less than 0, it means total is unknown.
func New(page, perPage int) *Pages {
	if perPage <= 0 {
		perPage = DefaultPageSize
	}
	if perPage > MaxPageSize {
		perPage = MaxPageSize
	}
	if page < 1 {
		page = 1
	}

	return &Pages{
		Pagination: Pagination{
			Page:    page,
			PerPage: perPage,
		},
	}
}

func (p *Pages) SetData(totalData int) {
	pageCount := -1
	if totalData >= 0 {
		pageCount = (totalData + p.PerPage - 1) / p.PerPage
	}

	p.PageCount = pageCount
	p.TotalCount = totalData
}

// NewFromRequest creates a Pages object using the query parameters found in the given HTTP request.
// count stands for the total number of items. Use -1 if this is unknown.
func NewFromRequest(c echo.Context) *Pages {
	page := parseInt(c.QueryParam(PageVar), 1)
	perPage := parseInt(c.QueryParam(PageSizeVar), DefaultPageSize)
	return New(page, perPage)
}

// parseInt parses a string into an integer. If parsing is failed, defaultValue will be returned.
func parseInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}

	if result, err := strconv.Atoi(value); err == nil {
		return result
	}

	return defaultValue
}

func SetPageLimit(page string, limit string) string {
	var offset int
	if page == "" {
		return " limit 100 offset 0"
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return ""
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return ""
	}
	if pageInt == 0 {
		limitInt = 0
	}

	offset = (pageInt - 1) * limitInt
	ret := fmt.Sprintf(" limit %d offset %d", limitInt, offset)
	return ret
}
