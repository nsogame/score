package score

import (
	"log"
	"net/http"
)

func Hwrapper(f func(http.ResponseWriter, *http.Request) error) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: don't throw away err
		err := f(w, r)
		if err != nil {
			log.Println("error:", err)
		}
	}
}
