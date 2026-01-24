package main

import (
	"fmt"
	"log"
	"net/http"
)

// go server to serve the public directory
func main() {
	fmt.Println("hello")
	// handler: listening for routes 
	// fileserver: serves file system
	http.Handle("/", http.FileServer(http.Dir("public")))
	// todo: move to .env files, one for dev, stage etc.
	const addr = ":8080"
	// TODO: what is the happening under the hood here
	err := http.ListenAndServe(addr, nil)
	if (err != nil) {
		log.Fatalf("Server failed, %v", err)
	}
}