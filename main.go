package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"pacgen/gener"

	"github.com/gin-gonic/gin"
)

func main() {
	var confFile string
	flag.StringVar(&confFile, "config", "config.toml", "configuration file")
	flag.Parse()

	var quitChannel = make(chan os.Signal, 1)
	signal.Notify(quitChannel, os.Interrupt, syscall.SIGTERM)

	g := gener.NewGener(confFile)
	stopTicker := make(chan struct{})
	g.WatchProxyMap(stopTicker)

	router := gin.Default()
	router.GET("/", g.GetPac)

	srv := &http.Server{
		Addr:    g.ListenAddr,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-quitChannel
	close(stopTicker)

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	log.Println("timeout of 5 seconds.")
	log.Println("Server exiting")
}
