package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"github.com/jwteeba/go-rest-api/helper"
	"github.com/jwteeba/go-rest-api/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// we created Book array
	var movies []models.Movie

	//Connection mongoDB with helper class
	collection := helper.ConnectDB()

	// bson.M{},  passed empty filter to get all data.
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	// Close the cursor once finished
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var movie models.Movie
		
		err := cur.Decode(&movie) 
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		movies = append(movies, movie)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(movies) 
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	var movie models.Movie
	// we get params with mux.
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])

	collection := helper.ConnectDB()

	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&movie)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(movie)
}

func main() {
	//Init Router
	r := mux.NewRouter()

  	// arrange our route
	r.HandleFunc("/api/movies", getMovies).Methods("GET")
	r.HandleFunc("/api/movies/{id}", getMovie).Methods("GET")
	// r.HandleFunc("/api/movies", createMovie).Methods("POST")
	// r.HandleFunc("/api/movies/{id}", updateMovie).Methods("PUT")
	// r.HandleFunc("/api/movies/{id}", deleteMovie).Methods("DELETE")

  	// set our port address
	log.Fatal(http.ListenAndServe(":8000", r))

}
