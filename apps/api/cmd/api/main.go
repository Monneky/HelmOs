package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"helmos/api/internal/config"
	"helmos/api/internal/db"
	"helmos/api/internal/handler"
	"helmos/api/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	if err := os.MkdirAll(cfg.DataDir, 0o750); err != nil {
		log.Fatalf("mkdir data dir: %v", err)
	}

	database, err := db.Open(cfg.DataDir)
	if err != nil {
		log.Fatalf("db open: %v", err)
	}
	defer db.Close()

	handler.SetDB(database)

	addr := ":" + strconv.Itoa(cfg.Port)
	srv := server.New(addr)

	done := make(chan struct{})
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
			log.Printf("shutdown: %v", err)
		}
		close(done)
	}()

	if err := srv.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server: %v", err)
	}
	<-done
}
