package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	port = flag.String("port", "8080", "port")
)

func main() {
	flag.Parse()
	var router = mux.NewRouter()
	var api = router.PathPrefix("/api").Subrouter()

	api.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	api.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("x-auth-token") != "admin" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			log.Println(r.RequestURI)
			next.ServeHTTP(w, r)
		})
	})

	var api1 = api.PathPrefix("/v1").Subrouter()
	api1.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	api1.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	})

	var api2 = api.PathPrefix("/v2").Subrouter()
	api2.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
	})
	api2.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	http.ListenAndServe(":"+*port, router)

}
