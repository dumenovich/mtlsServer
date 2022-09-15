package main

import (
	"crypto/tls"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Simple HTTP server
func echoMuxHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request connection: %s, path: %s", r.Proto, r.URL.Path[1:])

	vars := mux.Vars(r)
	value := vars["echo-string"]

	_, err := w.Write([]byte("Echo: " + value + "\n"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	w.WriteHeader(http.StatusCreated)
}

func main() {
	myMux := mux.NewRouter()

	myMux.HandleFunc("/v1/echo/{echo-string}", echoMuxHandler).Methods("PUT")

	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_128_CBC_SHA,
		},
	}

	srv := &http.Server{
		Addr:         ":8443",
		Handler:      myMux,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(server *http.Server, conn *tls.Conn, handler http.Handler), 0),
	}

	log.Printf("Go Backend HTPP 1/1: serving 8443 port, PUT /v1/echo/{}")

	//log.Fatal(http.ListenAndServe(":8443", myMux))
	log.Fatal(srv.ListenAndServeTLS("cert/server.crt", "cert/server.key"))
}
