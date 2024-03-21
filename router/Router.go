package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Router interface {
	Get(uri string, f func(w http.ResponseWriter, r *http.Request))
	Post(uri string, f func(w http.ResponseWriter, r *http.Request))
	Put(uri string, f func(w http.ResponseWriter, r *http.Request))
	Delete(uri string, f func(w http.ResponseWriter, r *http.Request))
	Serve(port string)
}

var (
	muxDisptcher = mux.NewRouter()
	api          = muxDisptcher.PathPrefix("/v1/api").Subrouter()
)

type MuxRouter struct {
	apiRoute   *mux.Router
	dispatcher *mux.Router
}

func NewRouter() Router {
	return &MuxRouter{apiRoute: api, dispatcher: muxDisptcher}
}

func (router *MuxRouter) Get(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	router.apiRoute.HandleFunc(uri, f).Methods("GET")
}

func (router *MuxRouter) Post(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	router.apiRoute.HandleFunc(uri, f).Methods("POST")
}

func (router *MuxRouter) Put(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	router.apiRoute.HandleFunc(uri, f).Methods("PUT")
}

func (router *MuxRouter) Delete(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	router.apiRoute.HandleFunc(uri, f).Methods("DELETE")
}

func (d *MuxRouter) Serve(port string) {
	fmt.Println("server is listening on port 5000")
	log.Fatal(http.ListenAndServe(port, d.dispatcher))
}
