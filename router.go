package score

import (
	"fmt"
	"net/http"
	"os"

	gorrilaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
}

func (score *ScoreServer) Handlers() (router http.Handler) {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/web/osu-submit-modular-selector.php", Hwrapper(score.SubmitModularHandler))
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// body, _ := ioutil.ReadAll(r.Body)
		fmt.Println("request", r.URL)
	})
	router = gorrilaHandlers.LoggingHandler(os.Stdout, r)
	return
}
