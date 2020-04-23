package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

const serverAddress = ":8080"

var logger *zap.SugaredLogger

func main() {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Errorf("cannot create zap.NewProduction, error: %w", err))
	}

	defer zapLogger.Sync()
	logger = zapLogger.Sugar()

	router := createRouter()
	http.Handle("/", router)
	startServer(router)
}

func startServer(router http.Handler) {
	server := &http.Server{Addr: serverAddress, Handler: router}
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		signal.Notify(sigint, syscall.SIGTERM)

		<-sigint

		logger.Info("shouting down the server")

		if err := server.Shutdown(context.Background()); err != nil {
			logger.Error("cannot shut down HTTP server, error: %v", err)
		}

		close(idleConnsClosed)
	}()

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logger.Error("HTTP server cannot ListenAndServe, error: %v", err)
	}

	logger.Info("HTPP server stopped")

	<-idleConnsClosed
}

func createRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health", healthHandler).Methods(http.MethodGet)
	r.HandleFunc("/echo", echoHandler).Methods(http.MethodGet)
	r.HandleFunc("/", echoHandler).Methods(http.MethodGet)
	return r
}

func echoHandler(w http.ResponseWriter, r *http.Request) {

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body := buf.String()

	logger.Infow("new request", "body", body)

	hostname, _ := os.Hostname()

	resp := map[string]interface{}{
		"body":     body,
		"method":   r.Method,
		"cookies":  r.Cookies(),
		"headers":  r.Header,
		"hostname": hostname}

	bts, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(bts); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("new health check request")
	w.WriteHeader(http.StatusOK)
}
