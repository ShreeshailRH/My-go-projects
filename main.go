package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Movie struct
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// Director struct
type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func main() {
	fmt.Println("Movies API with CRUD operations")

	// Initialize router
	r := mux.NewRouter()

	// Sample data
	movies = append(movies, Movie{ID: "1", Isbn: "3442334", Title: "One Movie", Director: &Director{FirstName: "Shree", LastName: "Huded"}})
	movies = append(movies, Movie{ID: "2", Isbn: "324223", Title: "Two Movie", Director: &Director{FirstName: "Ram", LastName: "Reddy"}})

	// Routes
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	// Start server
	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Handlers

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.NotFound(w, r)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = fmt.Sprintf("%d", len(movies)+1)
	movies = append(movies, movie)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	//set a json content type
	w.Header().Set("Content-Type", "application/json")
	// params
	params := mux.Vars(r)

	//loop over the moview
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)

			//update the movie with ID that you sent
			var updated Movie
			_ = json.NewDecoder(r.Body).Decode(&updated)
			updated.ID = params["id"]
			//add new movie to update
			movies = append(movies, updated)
			json.NewEncoder(w).Encode(updated)
			return
		}
	}
	http.NotFound(w, r)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}
