package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"frontendmasters.com/movies/data"
	"frontendmasters.com/movies/handlers"
	"frontendmasters.com/movies/logger"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// go server to serve the public directory
func main() {
	// Initialize logger
	logInstance := initializeLogger()

	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file was available")
	}
	fmt.Println("hello world")

	// Connect to DB
	dbConnStr := os.Getenv("DATABASE_URL")
	if dbConnStr == "" {
		log.Fatal("DATABASE_URL not set")
	}

	// here sql lib is trying to find a provider that can work w this database
	// there is no direct postgres support from go 
	// hence require importing pg library (driver)
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	// defer execution to the end 
	defer db.Close()

	// initlaise data repository for Movies
	// postgresql implementation to get data 
	movieRepo, err := data.NewMovieRepository(db, logInstance)
	if err != nil {
		log.Fatal("failed to initialise repository")
	}

	movieHandler := handlers.MovieHandler {}
	movieHandler.Storage = movieRepo
	movieHandler.Logger = logInstance

	// registers handler, for this path
	// to pass function as a argument, 
	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)
	http.HandleFunc("/api/movies/random", movieHandler.GetRandomMovies)
	// http.HandleFunc("/api/movies/", movieHandler.GetMovie)

	// handler for static files (frontend)
	// fileserver: serves file system
	// response to "GET http://localhost:8080/" request
	http.Handle("/", http.FileServer(http.Dir("public")))
	fmt.Println("serving the files")
	
	addr := os.Getenv("SERVER_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	logInstance.Info("Server starting on" + addr)
	// TODO: what is the happening under the hood here
	if err := http.ListenAndServe(addr, nil); (err != nil) {
		logInstance.Error("Server failed to start", err)
		log.Fatalf("Server failed, %v", err)
	}
}

func initializeLogger() *logger.Logger {
	logInstance, err := logger.NewLogger("movie-service.log")
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	} 
	return logInstance
}
