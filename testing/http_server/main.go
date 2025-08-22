package main

import (
	"flag"
	"fmt"
	"net/http"
)

var port = flag.Int("port", 8080, "port")

func main() {
	flag.Parse()
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"value":{"ready":true}}`))
	})
	http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", *port), nil)
}
