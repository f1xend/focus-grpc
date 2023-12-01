package main

import (
	"context"
	"fmt"
	"github.com/f1xend/focus-grpc/internal/illusionist/app"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	// err := godotenv.Load("F:\\GolangProjects\\focus-grpc\\.env")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("error loading .env file")
	}

	app.Start(ctx, fmt.Sprintf(":%s",
		os.Getenv("LESSON_ILLUSIONIST_HTTP_PORT")),
		fmt.Sprintf("%s:%s",
			os.Getenv("LESSON_ILLUSIONIST_FOCUS_URL"),
			os.Getenv("LESSON_ILLUSIONIST_FOCUS_PORT")),
	)
}
