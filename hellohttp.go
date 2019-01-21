package main
import (
	"fmt"
	"net/http"
)
func indexHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go Server")
	fmt.Fprintf(w, `<html>
		<body>
			Hello Gophers
		</body>
	</html>`,"\n")
}

func main() {
	http.HandleFunc("/", indexHandle)
	http.ListenAndServe(":8080", nil)
}
