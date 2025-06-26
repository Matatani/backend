package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"www.github.com/Maevlava/Matatani/backend/internal/config"
	"www.github.com/Maevlava/Matatani/backend/internal/ml"
	server "www.github.com/Maevlava/Matatani/backend/internal/server"
)

func newHTTPServer(cfg *config.APIConfig) *http.Server {
	return &http.Server{
		Addr:    ":" + cfg.HostPort,
		Handler: server.NewRouter(),
	}
}
func runHTTPServer(ctx context.Context, srv *http.Server) error {
	log.Printf("HTTP server listening on %s", srv.Addr)
	return srv.ListenAndServe()
}
func waitForShutdown(ctx context.Context, httpServer *http.Server) error {
	<-ctx.Done()
	log.Println("Shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Printf("HTTP shutdown error: %v", err)
	}

	log.Println("Servers shut down successfully.")

	return nil
}
func callMLService() error {
	var err error
	var conn *grpc.ClientConn
	var client ml.PredictorClient
	var ctx context.Context

	conn, err = grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client = ml.NewPredictorClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req := &ml.HelloRequest{
		Greeting: &ml.Greeting{
			Greeting: "Hello",
			Name:     "Matatani",
		},
		From: "Matatani",
	}

	resp, err := client.Hello(ctx, req)
	if err != nil {
		log.Fatalf("could not greet: %v\n", err)
	}
	log.Printf("Greeting: %s", resp.Greeting)

	return err
}

func main() {
	var err error
	var cfg *config.APIConfig

	cfg, err = config.Load()
	if err != nil {
		log.Printf("Error processing config: %s\nUsing OS env", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	grp, ctx := errgroup.WithContext(ctx)

	httpServer := newHTTPServer(cfg)

	grp.Go(func() error {
		return runHTTPServer(ctx, httpServer)
	})
	grp.Go(func() error {
		return waitForShutdown(ctx, httpServer)
	})

	go func() {
		err = callMLService()
		if err != nil {
			log.Printf("Error calling ML service: %v", err)
		}
	}()

	if err := grp.Wait(); err != nil {
		log.Printf("A server process failed: %v", err)
	}

	log.Println("Application shut down.")
}
