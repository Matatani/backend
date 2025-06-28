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
	"www.github.com/Maevlava/Matatani/backend/internal/predictor_service"
	server "www.github.com/Maevlava/Matatani/backend/internal/server"
)

func newHTTPServer(cfg *config.APIConfig, matataniServer *server.MatataniServer) *http.Server {
	addr := cfg.Host + ":" + cfg.Port
	log.Println("Binding to:", addr)
	return &http.Server{
		Addr:    addr,
		Handler: matataniServer.NewHTTPRouter(),
	}
}
func runHTTPServer(srv *http.Server) error {
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
func newMLClient(cfg *config.APIConfig) *predictor_service.PredictorClient {
	var client *predictor_service.PredictorClient

	conn, err := grpc.NewClient(cfg.MLHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect predictor_service service: %v", err)
	}
	defer conn.Close()

	return client
}

func main() {
	var err error
	var cfg *config.APIConfig

	cfg, err = config.Load()
	if err != nil {
		log.Printf("Error processing config: %s\nUsing OS env", err)
	}

	conn, err := grpc.NewClient(cfg.MLHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect predictor service: %v", err)
	}
	defer conn.Close()

	predictorClient := predictor_service.NewPredictorClient(conn)
	matataniServer := server.NewMatataniServer(cfg, predictorClient)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	grp, ctx := errgroup.WithContext(ctx)

	httpServer := newHTTPServer(cfg, matataniServer)

	grp.Go(func() error {
		return runHTTPServer(httpServer)
	})
	grp.Go(func() error {
		return waitForShutdown(ctx, httpServer)
	})

	if err := grp.Wait(); err != nil {
		log.Printf("A server process failed: %v", err)
	}

	log.Println("Application shut down.")
}
