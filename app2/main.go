package main

import (
	"bank2/pkg/handle"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println(handle.Print())
	http.HandleFunc("/bar", handle.FooHandler)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
