package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log"
	"net"
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
func newGRPCServer(cfg *config.APIConfig) *grpc.Server {
	srv := grpc.NewServer()
	ml.RegisterPredictorServer(srv, server.Predictor{})
	return srv
}
func runHTTPServer(ctx context.Context, srv *http.Server) error {
	log.Printf("HTTP server listening on %s", srv.Addr)
	return srv.ListenAndServe()
}
func runGRPCServer(ctx context.Context, address string, srv *grpc.Server) error {
	listener, _ := net.Listen("tcp", address)
	log.Printf("gRPC server listening on %s", listener.Addr())
	return srv.Serve(listener)
}
func waitForShutdown(ctx context.Context, httpServer *http.Server, grpcServer *grpc.Server) error {
	<-ctx.Done()
	log.Println("Shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Printf("HTTP shutdown error: %v", err)
	}

	grpcServer.GracefulStop()
	log.Println("Servers shut down successfully.")

	return nil
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Printf("Error processing config: %s\nUsing OS env", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	grp, ctx := errgroup.WithContext(ctx)

	httpServer := newHTTPServer(cfg)
	grpcServer := newGRPCServer(cfg)

	grp.Go(func() error {
		return runHTTPServer(ctx, httpServer)
	})
	grp.Go(func() error {
		return runGRPCServer(ctx, ":50051", grpcServer)
	})
	grp.Go(func() error {
		return waitForShutdown(ctx, httpServer, grpcServer)
	})

	if err := grp.Wait(); err != nil {
		log.Fatal(err)
	}
}
