package main

import (
	"fmt"
	"log"
	"net/http"
)

//Logger is first middleware
func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("in Logger")
		defer log.Println("Logger ended")
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()
		fmt.Println(r.Header)
		h.ServeHTTP(w, r)
	})
}

//LoggerType2 type2
func LoggerType2(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("in Logger")
		defer log.Println("Logger ended")
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()
		fmt.Println(r.Header)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

//HeaderValidator is second middleware
func HeaderValidator(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("in HeaderValidator")
		defer log.Println("HeaderValidator ended")
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()
		val := r.Header.Get("myheader")
		if val == "" {
			http.Error(w, "Error: no header found", http.StatusBadRequest)
			return
		}
		r.Header.Add("custom-response-header", "custom-value")
		h.ServeHTTP(w, r)
	})
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func main() {
	mux := http.NewServeMux()
	finalHandler := http.HandlerFunc(hello)
	mux.Handle("/hey", Logger(HeaderValidator(finalHandler)))
	log.Fatal(http.ListenAndServe(":8080", mux))
}
