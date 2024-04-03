package main 

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-Type","application/json")
	params := mux.Vars(r)
	for index,item := range movies {
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("content-Type","application/json")
	params := mux.Vars(r)
	for _,item := range movies {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return 
		}
	} 
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-Type","application/json")
    var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
}

func updateMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("content-Type","application/json")
	params := mux.Vars(r)

	for index,item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movies = append(movies,movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}

func main(){
	r := mux.NewRouter()

	movies = append(movies, Movie{ID:"1", Isbn:"435612", Title:"GOAT", Director:&Director{FirstName:"Venkat",LastName:"Prabhu"}})
	movies = append(movies, Movie{ID:"2", Isbn:"435452", Title:"T69", Director:&Director{FirstName:"Vetri",LastName:"Maran"}})
	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies",deleteMovies).Methods("DELETE")

	fmt.Printf("Starting Server at port 8080!\n")
	log.Fatal(http.ListenAndServe(":8080",r))
}

