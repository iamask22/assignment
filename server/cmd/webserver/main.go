package main

import (
	"context"
	"mta-hosting-optimizer/server/config/webserver"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	goCtx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()

	runServer(goCtx, cancel)
}

func runServer(goCtx context.Context, cancel context.CancelFunc) {
	srv := webserver.New(webserver.InitializeApplicationConfig)
	go func() {
		srv.ServeHTTP()
	}()

	waitForShutdownSignal(srv, goCtx, cancel)
}

func waitForShutdownSignal(srv *webserver.Server, goCtx context.Context, cancel context.CancelFunc) {
	var gracefulStop = make(chan os.Signal, 3)

	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	signal.Notify(gracefulStop, syscall.SIGQUIT)

	select {
	case <-gracefulStop:
		cancel()
		// if stop signal is received, wait for some time so that background workers get time to exit
		<-time.After(5 * time.Second)
	case <-goCtx.Done():
		// shutdown if context was cancelled by something else before shutdown signal
	}
	srv.Shutdown(goCtx)
}
