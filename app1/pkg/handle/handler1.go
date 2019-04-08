package handle

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//Print method which will print something
func Print() string {
	return "hello"
}

//FooHandler handler
func FooHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	endpointApp2 := os.Getenv("APP2")
	if endpointApp2 != "" {
		resp, err := http.Get(endpointApp2)
		if err != nil {
			fmt.Fprintf(w, "Error when calling service endpoint %s %s", endpointApp2, err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(w, "Error when reading the bosy %s", err)
		}
		fmt.Fprintf(w, string(body))
	} else {
		fmt.Fprintf(w, "{\"error\":\"endpoint not defined}\"")
	}

}

//Healthz test container health
func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}
