package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		mes := struct {
			Message string `json:"message"`
		}{
			Message: "pong",
		}
		b, err := json.Marshal(mes)
		if err != nil {
			panic(err)
		}
		w.Write(b)
	})

	http.ListenAndServe(":5000", r)
}
