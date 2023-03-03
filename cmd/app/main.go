package main

import (
	"log"
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
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func main() {
	conn, err := postgres.New(postgres.DBConfig{
		Host:     "localhost",
		Port:     5432,
		Database: "",
		User:     "",
		Password: "",
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

	var (
		repoStorage  = postgresStorage.New(conn.Pool, postgresStorage.Options{})
		_            = redisStorage.New(cache, redisStorage.Options{})
		ucContact    = useCaseContact.New(repoStorage, useCaseContact.Options{})
		ucGroup      = useCaseGroup.New(repoStorage, useCaseGroup.Options{})
		_            = deliveryGrpc.New(ucContact, ucGroup, deliveryGrpc.Options{})
		listenerHTTP = deliveryHttp.New(ucContact, ucGroup, deliveryHttp.Options{})
	)

	go func() {
		log.Printf("service started successfully on http port: %d", viper.GetUint("HTTP_PORT"))

		if err = listenerHTTP.Run(); err != nil {
			panic(err)
		}
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	<-signalCh
}
