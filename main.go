package main

import (
	"fmt"
	// "io"
	"log"
	"net/http"
)

func HelloServerHttp(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server listening on http.\n"))
	// fmt.Fprintf(w, "This is an example server.\n")
	// io.WriteString(w, "This is an example server.\n")
}

func HelloServerHttps(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Content-Type", "text/plain")
        w.Write([]byte("This is an example server listening on https.\n"))
        // fmt.Fprintf(w, "This is an example server.\n")
        // io.WriteString(w, "This is an example server.\n")
}

func main() {
	serverMuxHttp := http.NewServeMux()
	serverMuxHttp.HandleFunc("/", HelloServerHttp)

	serverMuxHttps := http.NewServeMux()
	serverMuxHttps.HandleFunc("/", HelloServerHttps)

	go func() {
		fmt.Println("Starting HTTP server")
		err := http.ListenAndServe(":8080", serverMuxHttp)
		if err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("Starting HTTPS server")
	err := http.ListenAndServeTLS(":8443", "server.crt", "server.key", serverMuxHttps)
	if err != nil {
		log.Fatal(err)
	}

}
