package rest

import (
	"net/http"
	"fmt"
	"time"
	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/schedules/{treatment}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		treatment := vars["treatment"]
		fmt.Fprintf(w, "Treatment %s \n", treatment)
	})
	r.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		treatment := vars["treatment"]
		fmt.Fprintf(w, "Treatment %s \n", treatment)
	})
	return r
	
}

func StartServer() {
	srv := &http.Server{
		Handler:      CreateRouter(),
		Addr:         "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
}