package main
import (
	"fmt"
	"net/http"
)

func main(){
	mux := http.ServeMux{}
	mux.Handle("/articles/", Handler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from /articles/")
	})
	mux.Handle("/users", Handler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from /users")
	})
	http.ListenAndServe(":3000", mux)
}
