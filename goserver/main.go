package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/zerogvt/idt/goserver/data"
	"github.com/zerogvt/idt/goserver/handlers"
)

func main() {
	bindaddr := ":8080"
	if len(os.Args) > 1 {
		bindaddr = os.Args[1]
	}
	l := log.New(os.Stdout, "products-api ", log.LstdFlags)
	udb := data.NewUserStore()
	data.FillWithMockUsers(udb)
	uh := handlers.NewUserHandler(udb, l)
	rtr := mux.NewRouter()
	usrRouter := rtr.PathPrefix("/user").Subrouter()
	usrRouter.Path("/{id}").Methods(http.MethodGet).HandlerFunc(uh.Get)
	usrRouter.Path("/{id}").Queries("name", "{name}").Methods(http.MethodPut).HandlerFunc(uh.Put)

	// create a new server
	srv := http.Server{
		Handler:      rtr,
		Addr:         bindaddr,          // configure the bind address
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Printf("Starting server on bind address %s", bindaddr)
		err := srv.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	//  catch kill or interupt and shutdown gracefully
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// wait till stopped by above signals
	sig := <-c
	l.Println("Got signal:", sig)

	// gracefully shutdown
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(ctx)
}
