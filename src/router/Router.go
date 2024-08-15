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
	RegisterRoute(method string, url string, f func(w http.ResponseWriter, r *http.Request))
	CreateSubRouter(prefix string) IRouter
	Serve(port string)
}

type MuxRouter struct {
	Router *mux.Router
}

func New() IRouter {
	return &MuxRouter{
		Router: mux.NewRouter(),
	}
}

func (r *MuxRouter) RegisterRoute(method string, url string, f func(w http.ResponseWriter, r *http.Request)) {
    r.Router.HandleFunc(url, f).Methods(method)

    // Debugging: Print out the registered routes
    r.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
        t, _ := route.GetPathTemplate()
        fmt.Println("Registered route:", t)
        return nil
    })
}

func (r *MuxRouter) CreateSubRouter(prefix string) IRouter {
    subRouter := r.Router.PathPrefix(prefix).Subrouter()
    return &MuxRouter{Router: subRouter}
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
