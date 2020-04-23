package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

func main() {
	router := createRouter()
	http.Handle("/", router)
	startServer(router)
}

func startServer(router http.Handler) {
	serverAddress := ":8080"
	server := &http.Server{Addr: serverAddress, Handler: router}

	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		signal.Notify(sigint, syscall.SIGTERM)

		<-sigint

		log.Printf("Shutdown: %v", true)
		// We received an interrupt signal, shut down.
		if err := server.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}

		close(idleConnsClosed)
	}()

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Printf("HTTP server ListenAndServe: %v", err)
	}

	log.Printf("Exiting: %v", true)

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
	hostname, _ := os.Hostname()

	resp := map[string]interface{}{}
	resp["cookies"] = r.Cookies()
	resp["headers"] = r.Header
	resp["hostname"] = hostname

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
	log.Printf("health")
	w.WriteHeader(http.StatusOK)
}
