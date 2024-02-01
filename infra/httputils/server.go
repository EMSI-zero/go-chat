package httputils

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func StartServer(httpserver *http.Server) (err error) {

	go func() {
		// zerolog.Infof("http server listening on %s", httpserver.Addr)
		log.Printf("http server listening on %s", httpserver.Addr)
		if err = httpserver.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// zerolog.Panic(err)
		}
	}()

	srv := httpserver
	chInterrupt := make(chan os.Signal, 1)
	signal.Notify(chInterrupt, os.Interrupt)
	chTerm := make(chan os.Signal, 1)
	signal.Notify(chTerm, syscall.SIGTERM)
	select {
	case <-chInterrupt:
		// zerolog.Info("server interrupted")
		stopServer(srv)
	case <-chTerm:
		// zerolog.Info("server received SIGTERM")
	}
	// zerolog.Info("done")

	return nil
}

func stopServer(srv *http.Server) {
	partStopped := make(chan struct{}, 20)
	go func() {
		const timeout = 5 * time.Second
		time.Sleep(timeout)
		partStopped <- struct{}{}
	}()
	go func() {
		// zerolog.Info(context.Background(), "stopping http server...")
		srvError := srv.Close()
		if srvError != nil {
			// zerolog.Info(context.Background(), "error: http server termination failed: %v", srvError)
		} else {
			// zerolog.Info(context.Background(), "stopped http server")
		}
		partStopped <- struct{}{}
	}()
	<-partStopped
}
