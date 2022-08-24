package main

import (
	"crypto/tls"
	"log"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr: ":8443",

		Handler: http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
				w.Write([]byte("hello, world\n"))
			},
		),

		TLSConfig: &tls.Config{
			NextProtos: []string{"istio", "istio-http/1.1", "istio-peer-exchange"},
		},
		TLSNextProto: map[string]func(*http.Server, *tls.Conn, http.Handler){
			"spdy/3": func(s *http.Server, conn *tls.Conn, h http.Handler) {
				buf := make([]byte, 1)
				if n, err := conn.Read(buf); err != nil {
					log.Panicf("%v|%v\n", n, err)
				}
			},
		},
	}

	err := server.ListenAndServeTLS("server.crt", "server.key")
	if err != nil {
		log.Fatal(err)
	}
}
