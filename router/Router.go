package router

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

type IRouter interface {
	Get(uri string, f func(w http.ResponseWriter, r *http.Request))
	Post(uri string, f func(w http.ResponseWriter, r *http.Request))
	Put(uri string, f func(w http.ResponseWriter, r *http.Request))
	Delete(uri string, f func(w http.ResponseWriter, r *http.Request))
	AddPrefix(prefix string) IRouter
	Serve(port string)
}

var (
	muxDisptcher = mux.NewRouter()
)

type MuxRouter struct {
	Router *mux.Router
}

func New(apiVersion string) IRouter {
	muxDisptcher.PathPrefix(fmt.Sprintf("/%s/api", apiVersion)).Subrouter()
	return &MuxRouter{Router: muxDisptcher}
}

func (r *MuxRouter) AddPrefix(prefix string) IRouter {
	subRouter := r.Router.PathPrefix(prefix).Subrouter()
	return &MuxRouter{Router: subRouter}
}

func (router *MuxRouter) Get(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	router.Router.HandleFunc(uri, f).Methods("GET")
}

func (router *MuxRouter) Post(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	router.Router.HandleFunc(uri, f).Methods("POST")
}

func (router *MuxRouter) Put(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	router.Router.HandleFunc(uri, f).Methods("PUT")
}

func (router *MuxRouter) Delete(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	router.Router.HandleFunc(uri, f).Methods("DELETE")
}

func (router *MuxRouter) Serve(port string) {
	srv := &http.Server{
		Addr:    port,
		Handler: router.Router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", port, err)
		}
	}()

	fmt.Printf("Server is listening on port %s\n", port)

	// Graceful shutdown logic
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	fmt.Println("Shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	fmt.Println("Server exiting")
}
