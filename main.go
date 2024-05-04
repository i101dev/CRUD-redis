package main

import (
	"log"
	"net/http"
)

func main() {

	// `&` indicates that we are taking the memory address, rather than the value
	server := &http.Server{
		Addr:    ":5000",
		Handler: http.HandlerFunc(basicHandler),
	}

	if err := server.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}

// Parameters are in reverse order compared to Nodejs
// `ResponseWriter` equivalent to [express] ->`res/response`
func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Say what again"))
}
