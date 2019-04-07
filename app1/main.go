package main

import (
	"bank/pkg/handle"
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	fmt.Println(handle.Print())
	http.HandleFunc("/foo", handle.FooHandler)
	http.HandleFunc("/healthz", handle.Healthz)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8081", nil))
}
