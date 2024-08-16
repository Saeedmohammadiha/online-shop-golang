package router

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/OnlineShop/internal/pkg/logger"
	"github.com/gorilla/mux"
)

type IRouter interface {
	RegisterRoute(method string, url string, f func(w http.ResponseWriter, r *http.Request))
	CreateSubRouter(prefix string) IRouter
	Use(middleware func(http.Handler) http.Handler)
	Serve(port string)
}

type MuxRouter struct {
	Router *mux.Router
	log    logger.Ilogger
}

func New(log logger.Ilogger) IRouter {
	log.Info("router is created")
	return &MuxRouter{
		Router: mux.NewRouter(),
		log:    log,
	}
}

func (r *MuxRouter) RegisterRoute(method string, url string, f func(w http.ResponseWriter, r *http.Request)) {
	r.Router.HandleFunc(url, f).Methods(method)
	r.log.Info("a route registered",
		"method", method,
		"url", url,
	)
	// Debugging: Print out the registered routes
	r.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, _ := route.GetPathTemplate()
		r.log.Debug("Registered route:", t)
		return nil
	})
}

func (r *MuxRouter) CreateSubRouter(prefix string) IRouter {
	 r.log.Info("a subRouter is created", "rout:", prefix)
	subRouter := r.Router.PathPrefix(prefix).Subrouter()
	return &MuxRouter{Router: subRouter, log: r.log}
}

func (r *MuxRouter) Use(middleware func(http.Handler) http.Handler) {
	r.log.Info("added a middleware to router", "middleware", middleware)
	r.Router.Use(middleware)
}

func (r *MuxRouter) Serve(port string) {
	srv := &http.Server{
		Addr:    port,
		Handler: r.Router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			r.log.Fatalf("Could not listen on %s: %v\n", port, err)
		}
	}()

	fmt.Printf("Server is listening on port %s\n", port)
	r.log.Infof("Server is listening on port %s\n", port)

	// Graceful shutdown logic
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	r.log.Info("Shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		r.log.Fatalf("Server forced to shutdown: %v", err)
	}

	r.log.Info("Server exiting")
}
