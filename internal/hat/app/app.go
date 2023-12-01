package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"os/signal"
	"syscall"

	"github.com/f1xend/focus-grpc/internal/hat/controller"
	"github.com/f1xend/focus-grpc/pkg/hat"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

// Start запускает сервис на указанном адресе и обрабатывает сигналы завершения.
// Принимает контекст и адрес в качестве аргументов.
func Start(ctx context.Context, addr string) {
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	gr, ctx := errgroup.WithContext(ctx)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println(err)
	}

	server := grpc.NewServer()
	hat.RegisterHatServer(server, new(controller.Controller))

	gr.Go(func() error {
		log.Printf("start service (%s)\n", addr)
		if err = server.Serve(lis); err != nil {
			log.Println(fmt.Errorf("error service (%s): %w", addr, err))
		}
		return nil
	})

	gr.Go(func() error {
		<-ctx.Done()
		server.GracefulStop()
		log.Printf("stop service (%s)\n", addr)
		return nil
	})

	if err = gr.Wait(); err != nil {
		log.Println(fmt.Errorf("error service (%s): %w", addr, err))
	}
}
