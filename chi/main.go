package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	r.Get("/random", answerRandom)

	r.Post("/random", errorPage)
	http.ListenAndServe(":8085", r)
}

func answerRandom(w http.ResponseWriter, r *http.Request) {
	req := r.URL.Query()

	if len(req) == 2 {
		min, err := strconv.Atoi(req["min"][0])
		if err != nil {
			log.Println("impossible convert to integer")
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		max, err := strconv.Atoi(req["max"][0])
		if err != nil {
			log.Println("impossible convert to integer")
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		if min >= max {
			http.Error(w, "'min' can't be greater than 'max' or equal to 'max'", http.StatusBadRequest)
			return
		}
		result := strconv.Itoa(rand.Intn(max-min+1) + min)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", result)
	} else {
		http.Error(w, "Wrong amount of params", http.StatusBadRequest)
	}
}

func errorPage(w http.ResponseWriter, req *http.Request) {
	http.Error(w, "Bad request", http.StatusBadRequest)
}
