package main

import (
	"fmt"
	"log"
	"net/http"

	"frontendmasters.com/movies/handlers"
)

// go server to serve the public directory
func main() {
	fmt.Println("hello")

	movieHandler := handlers.MovieHandler {}
	// registers handler, for this path
	// to pass function as a argument, 
	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)
	// handler: listening for routes 
	// handler for static files (frontend)
	// fileserver: serves file system
	http.Handle("/", http.FileServer(http.Dir("public")))
	fmt.Println("serving the files")
	
	// todo: move to .env files, one for dev, stage etc.
	const addr = ":8080"
	// TODO: what is the happening under the hood here
	err := http.ListenAndServe(addr, nil)
	if (err != nil) {
		log.Fatalf("Server failed, %v", err)
	}


}