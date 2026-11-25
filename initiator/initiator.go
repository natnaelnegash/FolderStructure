package initiator

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"folderstructure/internal/constant/model/persistencedb"
	"folderstructure/internal/handler/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Initiate() {
	ctx := context.Background()

	log, err := zap.NewProduction()
	if err != nil {
		log.Fatal("unable to start logger")
	}

	log.Info("initializing config")
	configName := "config"
	if os.Getenv("CONFIG_NAME") != "" {
		configName = os.Getenv("CONFIG_NAME")
	}

	err = InitConfig(Config{
		Names:  []string{configName},
		Path:   "config",
		Logger: log,
	})
	if err != nil {
		log.Fatal("unable to start config", zap.Error(err))
	}
	log.Info("initializing config completed")

	logger := InitLogger()
	log.Info("initializing databases")
	pgxPool, mongoClient, mongoDB := initDatabases(logger)

	log.Info("initializing persistence layer")
	persistenceDB := persistencedb.New(pgxPool, mongoClient, mongoDB, logger)

	redisPool := initRedis(viper.GetString("redis.url"), logger)
	log.Info("redis connection initialized")

	wp := initWorkerPool(viper.GetInt("workerpool.max_workers"), viper.GetInt("workerpool.task_buffer"))

	persistence := initPersistence(&persistenceDB, logger, redisPool)
	logger.Info(ctx, "done initializing persistence layer")
	logger.Info(ctx, "initializing kafka")

	logger.Info(ctx, "initializing module layer")
	service := initService(persistence, logger, *wp)
	logger.Info(ctx, "done initializing service layer")

	logger.Info(ctx, "initializing transport layer")
	transport := initHandler(service, logger)
	logger.Info(ctx, "done initializing transport layer")

	r := chi.NewRouter()

	r.Use(middleware.Logger(logger))
	r.Use(middleware.CORS())
	r.Use(middleware.ErrorHandler)

	initRoute(r, transport, service, logger)
	logger.Info(ctx, "done initializing route")

	srv := &http.Server{
		Addr:              fmt.Sprintf("%s:%d", viper.GetString("app.host"), viper.GetInt("app.port")),
		Handler:           r,
		ReadHeaderTimeout: viper.GetDuration("app.timeout"),
		IdleTimeout:       30 * time.Minute,
	}
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint

		logger.Info(ctx, "shutting down HTTP server...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			logger.Error(ctx, "HTTP server shutdown error", zap.Error(err))
		}
	}()

	logger.Info(ctx, fmt.Sprintf("http server listening on %s", srv.Addr))
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal(ctx, fmt.Sprintf("Could not start HTTP server: %s", err))
	}
}
