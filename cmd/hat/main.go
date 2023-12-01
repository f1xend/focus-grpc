package main

import (
	"context"
	"fmt"
	"os"

	"github.com/f1xend/focus-grpc/internal/hat/app"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("error loading .env file")
	}
	app.Start(ctx, fmt.Sprintf(":%s", os.Getenv("LESSON_HAT_GRPC_PORT")))
}
