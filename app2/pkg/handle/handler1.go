package handle

import (
	"fmt"
	"net/http"
	"os"
)

//Print method which will print something
func Print() string {
	return "hello"
}

//BarHandler handler
func BarHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	endpointApp1 := os.Getenv("APP1")
	json := "{'id':'1234','name':'Victor','endpointApp1':" + endpointApp1 + "}"
	fmt.Fprintf(w, json)

}

//Healthz test container health
func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}
