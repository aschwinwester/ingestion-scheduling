package rest

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/schedules/{treatment}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		treatment := vars["treatment"]
		fmt.Fprintf(w, "Treatment %s \n", treatment)
	})
	return r
	
}

func StartServer() {
	http.ListenAndServe(":8080", CreateRouter())
}