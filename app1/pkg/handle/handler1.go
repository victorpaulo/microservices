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

//FooHandler handler
func FooHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "Hi there APP1, I love %s!<br/>", r.URL.Path[1:])
	fmt.Fprintf(w, "Call <a href=%s>app2</a>", os.Getenv("APP2"))
}

//Healthz test container health
func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}
