package main

import (
	"context"
	"log"
	"github.com/rtmelsov/simple-db-psql/internal/database"
	"fmt"
)

func main() {
	pool, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer pool.Close()

	var greeting string
	err = pool.QueryRow(context.Background(), "SELECT 'Hello world!'").Scan(&greeting)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("greeting", greeting)
}
