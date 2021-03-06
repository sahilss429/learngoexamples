package main
import (
	"fmt"
	"net/http"
	"time"
)

type UptimeHandler struct {
	Started time.Time
}

func NewUptimeHandler() UptimeHandler{
	return UptimeHandler{ Started: time.Now() }
}

func (h UptimeHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, fmt.Sprintf("Current Uptime: %s", time.Since(h.Started)),)
}

func main() {
	http.Handle("/", NewUptimeHandler())
	http.ListenAndServe(":8080", nil)
}
