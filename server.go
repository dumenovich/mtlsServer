package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func echoMuxHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request connection: %s, path: %s", r.Proto, r.URL.Path[1:])

	vars := mux.Vars(r)
	value := vars["echo-string"]

	_, err := w.Write([]byte("Echo: " + value + "\n"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func main() {
	myMux := mux.NewRouter()

	myMux.HandleFunc("/v1/echo/{echo-string}", echoMuxHandler).Methods("PUT")

	log.Printf("Go Backend HTPP 1/1: serving 8443 port, PUT /v1/echo/{}")
	log.Fatal(http.ListenAndServe(":8443", myMux))
	//log.Fatal(http.ListenAndServeTLS(":8443", "pki/cert.pem", "pki/key.pem", nil))
}
