package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/hellow", basicHandler)

	// `&` indicates that we are taking the memory address, rather than the value
	server := &http.Server{
		Addr:    ":5000",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println()
		fmt.Print("*** >>> Failed to launch server - Error:")
		fmt.Println()
		log.Panic(err)
		fmt.Println()
	}
}

// Parameters are in reverse order compared to Nodejs
// `ResponseWriter` equivalent to [express] ->`res/response`
func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Say what again"))
}
