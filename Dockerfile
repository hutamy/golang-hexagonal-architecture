# Stage 1: Build the Go application
FROM golang:1.23.3-alpine3.19 AS builder

WORKDIR /app

COPY . .

RUN apk update && apk upgrade --no-cache && \
  apk add --no-cache git libcrypto3 libssl3

RUN go mod tidy && \
  go mod download && \
  go build -o golang-hexagonal-architecture


# Stage 2: Create the final lightweight image
FROM alpine:3.19


WORKDIR /app

# Copy the built binary and Datadog initialization script from the builder stage
COPY --from=builder /app/golang-hexagonal-architecture /app/golang-hexagonal-architecture

RUN touch .env

# Create a non-root user called 'appuser'
RUN adduser -D -g '' appuser

# Change ownership of the binary to the non-root user
RUN chown appuser:appuser /app/golang-hexagonal-architecture

# Switch to the non-root user
USER appuser

# Expose the port
EXPOSE 5531

# CMD ./golang-hexagonal-architecture 
CMD ["/app/golang-hexagonal-architecture"]

