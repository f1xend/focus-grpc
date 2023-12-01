package app

import (
	"context"
	"fmt"
	"github.com/f1xend/focus-grpc/internal/illusionist/repository"
	"github.com/f1xend/focus-grpc/internal/illusionist/usecase"
	"github.com/f1xend/focus-grpc/pkg/hat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/f1xend/focus-grpc/internal/illusionist/controller"
	"golang.org/x/sync/errgroup"
)

func Start(ctx context.Context, addr, hatAddr string) {
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	l := log.New(os.Stdout, "", 0)

	repo, err := initShowRepository(hatAddr)
	if err != nil {
		log.Println(fmt.Errorf("error init repository: %w", err))
		return
	}

	showRepo := repository.NewRabbit(l, repo)
	showUseCase := usecase.NewShow(showRepo, l)
	c := controller.NewController(l, showUseCase)
	mux := route(c)

	//logHandler := logMiddleware(l)
	httpServer := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		log.Printf("старт сервиса (%s)\n", addr)
		return httpServer.ListenAndServe()
	})

	g.Go(func() error {
		<-ctx.Done()
		err := httpServer.Shutdown(context.Background())
		if err != nil {
			return err
		}
		log.Printf("завершение работы сервиса (%s)\n", addr)
		return nil
	})

	if err := g.Wait(); err != nil {
		if err == http.ErrServerClosed {
			return
		}
		log.Println(fmt.Errorf("ошибка сервиса (%s): %w", addr, err))
	}
}

func initShowRepository(addr string) (hat.HatClient, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Println(fmt.Errorf("error connect to (%s):%v", addr, err))
		return nil, err
	}
	return hat.NewHatClient(conn), err
}
