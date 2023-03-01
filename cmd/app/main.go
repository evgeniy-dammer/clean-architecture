package main

import (
	"fmt"

	"github.com/evgeniy-dammer/clean-architecture/pkg/store/postgres"
)

func main() {
	conn, err := postgres.NewPostgres(postgres.DBConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "emenu",
		Password: "emenu",
		Database: "emenu",
		SSLMode:  "disable",
	})
	if err != nil {
		panic(err)
	}

	defer conn.Pool.Close()

	fmt.Println(conn.Pool.Stat())
}
