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
	fmt.Fprintf(w, "Hi there APP2, I love %s!", r.URL.Path[1:])
	fmt.Fprintf(w, "Call app1 <a href=%s></a>", os.Getenv("APP1"))
}
