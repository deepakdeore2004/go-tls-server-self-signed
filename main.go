package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
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
		//TLSNextProto: map[string]func(*http.Server, *tls.Conn, http.Handler){
		// "istio": handleIstio,
		// "istio-http/1.1": handleIstio,
		// "istio-peer-exchange": handleIstio,
		//},
		TLSNextProto: map[string]func(*http.Server, *tls.Conn, http.Handler){
			"spdy/3": func(s *http.Server, conn *tls.Conn, h http.Handler) {
				fmt.Println("In handler spdy/3")
				buf := make([]byte, 1)
				if n, err := conn.Read(buf); err != nil {
					log.Panicf("%v|%v\n", n, err)
				}
			},
			"istio": func(s *http.Server, conn *tls.Conn, h http.Handler) {
				fmt.Println("In handler istio")
				buf := make([]byte, 1)
				if n, err := conn.Read(buf); err != nil {
					log.Panicf("%v|%v\n", n, err)
				}
				fmt.Println("Trying to write in")
				wtr := bufio.NewWriter(conn)
				_, err := wtr.WriteString("\nRequested protocol: istio\n\n")
				fmt.Println("Trying to write in --1 ")
				if err != nil {
					fmt.Printf("error during istio client writing: %v\n", err)
				}
				wtr.Flush()
			},
			"istio-http/1.1": func(s *http.Server, conn *tls.Conn, h http.Handler) {
				fmt.Println("In handler istio-http/1.1")
				buf := make([]byte, 1)
				if n, err := conn.Read(buf); err != nil {
					log.Panicf("%v|%v\n", n, err)
				}
				fmt.Println("Trying to write in")
				wtr := bufio.NewWriter(conn)
				_, err := wtr.WriteString("\nRequested protocol: istio-http/1.1\n\n")
				fmt.Println("Trying to write in --1 ")
				if err != nil {
					fmt.Printf("error during istio client writing: %v\n", err)
				}
				wtr.Flush()

			},
			"istio-peer-exchange": func(s *http.Server, conn *tls.Conn, h http.Handler) {
				buf := make([]byte, 1)
				fmt.Println("In handler istio-peer-exchange")
				if n, err := conn.Read(buf); err != nil {
					log.Println("hi")
					log.Panicf("%v|%v\n", n, err)
				}
				fmt.Println("Trying to write in")
				wtr := bufio.NewWriter(conn)
				_, err := wtr.WriteString("\nRequested protocol: istio-peer-exchange\n\n")
				fmt.Println("Trying to write in --1 ")
				if err != nil {
					fmt.Printf("error during istio client writing: %v\n", err)
				}
				wtr.Flush()
			},
		},
	}
	err := server.ListenAndServeTLS("server.crt", "server.key")
	if err != nil {
		log.Fatal(err)
	}
}
