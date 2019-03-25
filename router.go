package score

import (
	"net/http"
	"os"

	gorrilaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
}

func handler() (router http.Handler) {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	router = gorrilaHandlers.LoggingHandler(os.Stdout, r)
	return
}
