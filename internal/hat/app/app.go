package app

import (
	"context"
	"fmt"
	"github.com/f1xend/focus-grpc/internal/hat/repository"
	"github.com/f1xend/focus-grpc/internal/hat/usecase"
	"github.com/f1xend/focus-grpc/pkg/postgres"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/f1xend/focus-grpc/internal/hat/controller"
	"github.com/f1xend/focus-grpc/pkg/hat"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

// Start запускает сервис на указанном адресе и обрабатывает сигналы завершения.
// Принимает контекст и адрес в качестве аргументов.
func Start(ctx context.Context, addr string, pgConn postgres.PgConn) {
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	l := log.New(os.Stdout, "", 0)

	gr, ctx := errgroup.WithContext(ctx)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		l.Println(err)
	}

	server := grpc.NewServer()

	animalRepo, err := initAnimalRepo(ctx, pgConn, l)
	if err != nil {
		l.Println(fmt.Errorf("error init repo (%s): %w", addr, err))
	}
	rabbitUsecase := usecase.NewRabbit(animalRepo, l)
	controller := controller.NewController(rabbitUsecase, l)
	hat.RegisterHatServer(server, controller)

	gr.Go(func() error {
		l.Printf("start service (%s)\n", addr)
		if err = server.Serve(lis); err != nil {
			l.Println(fmt.Errorf("error service (%s): %w", addr, err))
		}
		return nil
	})

	gr.Go(func() error {
		<-ctx.Done()
		server.GracefulStop()
		l.Printf("stop service (%s)\n", addr)
		return nil
	})

	if err = gr.Wait(); err != nil {
		l.Println(fmt.Errorf("error service (%s): %w", addr, err))
	}
}

func initAnimalRepo(ctx context.Context, conn postgres.PgConn, l *log.Logger) (repository.Rabbit, error) {
	db, err := postgres.NewPg(ctx, conn, l)
	if err != nil {
		l.Println(fmt.Errorf("error init db: %w", err))
		return repository.Rabbit{}, nil
	}
	return repository.NewRabbit(db, l), nil
}
