package main

import (
	"context"
	"fmt"
	"github.com/f1xend/focus-grpc/pkg/postgres"
	"log"
	"os"

	"github.com/f1xend/focus-grpc/internal/hat/app"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("error loading .env file")
	}

	pgConf := postgres.PgConn{
		Host:     os.Getenv("LESSON_HAT_PG_DB_HOST"),
		Port:     os.Getenv("LESSON_HAT_PG_DB_PORT"),
		User:     os.Getenv("LESSON_HAT_PG_DB_USER"),
		Password: os.Getenv("LESSON_HAT_PG_DB_PASSWORD"),
		Db:       os.Getenv("LESSON_HAT_PG_DB_NAME"),
	}

	app.Start(ctx, fmt.Sprintf(":%s", os.Getenv("LESSON_HAT_GRPC_PORT")), pgConf)
}
