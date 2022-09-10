package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

const serverAddress = ":8081"

func main() {
	http.HandleFunc("/random", answerRandom)
	// http.HandleFunc("/random", answerRandom)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK) // 200
		fmt.Fprintln(w, r)
	})

	http.ListenAndServe(serverAddress, nil)
}

// return random number between specified min and max
func answerRandom(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet: // "GET"
		r := req.URL.Query()
		if len(r) == 2 {
			min, err := strconv.Atoi(r["min"][0])
			if err != nil {
				log.Println("impossible convert to integer")
				http.Error(w, "Bad request", http.StatusBadRequest)
				return
			}
			max, err := strconv.Atoi(r["max"][0])
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
			w.Write([]byte("<h1>" + result + "</h1>"))
		} else {
			http.Error(w, "Wrong amount of params", http.StatusBadRequest)
		}
	case http.MethodPost:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Sorry, only GET is supported.")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Sorry, only GET method is supported.")
	}
}
