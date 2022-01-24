package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"phamacy/handlerStuff"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// Init router
	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/drugs", handlerStuff.GetDrugs).Methods("GET")
	r.HandleFunc("/drug/{id}", handlerStuff.GetDrug).Methods("GET")
	r.HandleFunc("/drugs", handlerStuff.CreateDrug).Methods("POST")
	r.HandleFunc("/drugs/{id}", handlerStuff.UpdateDrug).Methods("PUT")
	r.HandleFunc("/drugs/{id}", handlerStuff.DeleteDrug).Methods("DELETE")

	s := &http.Server{
		Addr:              ":9001",
		Handler:           r,
		WriteTimeout:      3 * time.Second,
		ReadTimeout:       3 * time.Second,
		ReadHeaderTimeout: 3 * time.Second,
		IdleTimeout:       5 * time.Second,
		// ErrorLog:          err,
	}

	go func() {
		s.ListenAndServe()
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Println("Gracefully shutting down server due to ", sig)

	tx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	s.Shutdown(tx)
}
