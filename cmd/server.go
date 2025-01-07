package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hutamy/golang-hexagonal-architecture/config"

	iregistry "github.com/hutamy/golang-hexagonal-architecture/internal/adapter/inbound/registry"
	"github.com/hutamy/golang-hexagonal-architecture/internal/adapter/inbound/rest"
	"github.com/hutamy/golang-hexagonal-architecture/internal/adapter/outbound/datastore"
	mregistry "github.com/hutamy/golang-hexagonal-architecture/internal/adapter/outbound/repository/registry"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	echoDatadog "gopkg.in/DataDog/dd-trace-go.v1/contrib/labstack/echo.v4"
	dd_logrus "gopkg.in/DataDog/dd-trace-go.v1/contrib/sirupsen/logrus"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
)

func RunServer() {
	cfg := config.GetConfig()
	e := echo.New()
	ctx := context.Background()

	tracer.Start(tracer.WithDebugMode(false))
	defer tracer.Stop()

	err := profiler.Start(
		profiler.WithService(cfg.DDService),
		profiler.WithEnv(cfg.DDEnv),
		profiler.WithProfileTypes(
			profiler.CPUProfile,
			profiler.HeapProfile,
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer profiler.Stop()

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.AddHook(&dd_logrus.DDContextLogHook{})

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Use(echoDatadog.Middleware())

	// Init mongodb connection
	mUser := cfg.MongoUser
	mPass := cfg.MongoPassword
	mHost := cfg.MongoHost
	mName := cfg.MongoName

	mdb = InitMongoModule(mUser, mPass, mHost, mName)
	defer datastore.CloseMongo(mdb)

	// Init postgresdb connection
	pHost := cfg.PostgresHost
	pPort := cfg.PostgresPort
	pUser := cfg.PostgresUsername
	pPass := cfg.PostgresPassword
	pDatabase := cfg.PostgresDatabase

	db = InitPostgresModule(pHost, pPort, pUser, pPass, pDatabase)
	defer db.Close()

	// Init postgresdb replica connection
	prHost := cfg.PostgresHostReplica
	prPort := cfg.PostgresPortReplica
	prUser := cfg.PostgresUsernameReplica
	prPass := cfg.PostgresPasswordReplica
	prDatabase := cfg.PostgresDatabaseReplica

	rdb = InitPostgresModule(prHost, prPort, prUser, prPass, prDatabase)
	defer rdb.Close()

	// Init redis connection
	rHost := cfg.RedisHost
	rPort := cfg.RedisPort
	rPassword := cfg.RedisPassword

	cache = InitRedisModule(ctx, rHost, rPassword, rPort)
	defer cache.Close()

	// Init repository registry
	repositoryRegistry := mregistry.NewRepositoryRegistry(mdb, db, rdb, cache)

	// Init service registry
	serviceRegistry := iregistry.NewServiceRegistry(repositoryRegistry)

	rest.Apply(e, serviceRegistry)
	go func() {
		e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.Port)))
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	signal := <-c
	log.Fatalf("process killed with signal: %v\n", signal.String())
}
