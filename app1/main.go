package main

import (
	"bank/pkg/handle"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println(handle.Print())
	http.HandleFunc("/foo", handle.FooHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}