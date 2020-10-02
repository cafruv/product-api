package main

import (
<<<<<<< HEAD
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
=======
	"log"
	"net/http"
	"os"
>>>>>>> 5180714a767b170a1b2807c7520d6314ac5d5251

	"github.com/cafruv/go-api/handlers"
)

func main() {
	l := log.New(os.Stdout, "go-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)
	sm := http.NewServeMux()
<<<<<<< HEAD

	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	// create a new server
	s := http.Server{
		Addr:         "localhost:9090",  // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)

=======

	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	http.ListenAndServe(":9000", sm)
>>>>>>> 5180714a767b170a1b2807c7520d6314ac5d5251
}
