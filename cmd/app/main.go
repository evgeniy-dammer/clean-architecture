package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	deliveryGrpc "github.com/evgeniy-dammer/clean-architecture/internal/delivery/grpc"
	deliveryHttp "github.com/evgeniy-dammer/clean-architecture/internal/delivery/http"
	postgresStorage "github.com/evgeniy-dammer/clean-architecture/internal/repository/storage/postgres"
	redisStorage "github.com/evgeniy-dammer/clean-architecture/internal/repository/storage/redis"
	useCaseContact "github.com/evgeniy-dammer/clean-architecture/internal/usecase/contact"
	useCaseGroup "github.com/evgeniy-dammer/clean-architecture/internal/usecase/group"
	"github.com/evgeniy-dammer/clean-architecture/pkg/store/postgres"
	redisCache "github.com/evgeniy-dammer/clean-architecture/pkg/store/redis"
	"github.com/evgeniy-dammer/clean-architecture/pkg/tracing"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/context"
	log "github.com/evgeniy-dammer/clean-architecture/pkg/type/logger"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	conn, err := postgres.New(postgres.DBConfig{
		Host:     "localhost",
		Port:     5432,
		Database: "clean",
		User:     "clean",
		Password: "clean",
		SSLMode:  "disable",
	})
	if err != nil {
		panic(err)
	}

	defer conn.Pool.Close()

	cache, err := redisCache.NewRedisCache(redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	if err != nil {
		panic(err)
	}

	closer, err := tracing.New(context.Empty())
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = closer.Close(); err != nil {
			log.Error(err)
		}
	}()

	repoStorage, err := postgresStorage.New(conn.Pool, postgresStorage.Options{})
	if err != nil {
		panic(err)
	}

	var (
		_            = redisStorage.New(cache, redisStorage.Options{})
		ucContact    = useCaseContact.New(repoStorage, useCaseContact.Options{})
		ucGroup      = useCaseGroup.New(repoStorage, useCaseGroup.Options{})
		_            = deliveryGrpc.New(ucContact, ucGroup, deliveryGrpc.Options{})
		listenerHTTP = deliveryHttp.New(ucContact, ucGroup, deliveryHttp.Options{})
	)

	go func() {
		fmt.Printf("service started successfully on http port: %d", viper.GetUint("HTTP_PORT"))

		if err = listenerHTTP.Run(); err != nil {
			panic(err)
		}
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	<-signalCh
}
