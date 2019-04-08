package main

import (
	"bank2/pkg/handle"
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	fmt.Println(handle.Print())
	http.HandleFunc("/bar", handle.BarHandler)
	http.HandleFunc("/healthz", handle.Healthz)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8082", nil))
}
