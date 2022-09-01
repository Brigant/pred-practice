package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/random/", answerRandom)
	http.HandleFunc("/random", answerRandom)

	http.ListenAndServe(":8081", nil)
}

// return random number between specified min and max
func answerRandom(w http.ResponseWriter, req *http.Request) {
	log.Println(req.URL.Path)
	log.Println(req.Method)
	log.Println(req.URL.Query())

	if req.URL.Path != "/random/" {
		http.Error(w, "404 not found or try /random/ with parameters", http.StatusNotFound)
		return
	}

	switch req.Method {
	case http.MethodGet:
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
