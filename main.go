package main

import (
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const addr = ":3000"

func main() {
	r := mux.NewRouter()

	{
		p := int64(0)
		r.HandleFunc("/count/", func(http.ResponseWriter, *http.Request) {
			atomic.AddInt64(&p, 1)
		}).Methods("PUT")
		r.HandleFunc("/count/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(fmt.Sprintln(p)))
		}).Methods("GET")
		// Uncomment the following line to enable CORS.
		// r.HandleFunc("/count/", func(w http.ResponseWriter, r *http.Request) {
		// 	w.WriteHeader(http.StatusMethodNotAllowed)
		// })
	}

	r.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.Method, r.URL.Path)
			h.ServeHTTP(w, r)
		})
	})
	r.Use(handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"PUT", "GET", "OPTIONS"}),
	))

	fmt.Println("listening on", addr)
	http.ListenAndServe(addr, r)
}
