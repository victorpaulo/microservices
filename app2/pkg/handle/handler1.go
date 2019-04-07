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
	fmt.Fprintf(w, "Hi there APP2, I love %s!<br/>", r.URL.Path[1:])
	fmt.Fprintf(w, "Call <a href=%s>app1</a>", os.Getenv("APP1"))
}
