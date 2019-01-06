package main

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/kazukousen/gorea/static"
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

	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "static")
	static.FileServer(r, "/views", http.Dir(filesDir))

	http.ListenAndServe(":5000", r)
}
