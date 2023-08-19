package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongoapi/model"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://test:saksham@cluster0.izu8bqz.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const collectionName = "watchlist"

// collection object/instance
var collection *mongo.Collection // original reference to collection

// connect to MongoDB
func init() {
	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions) // context.BACKGROUND() keeps connection alive in the background
	if err != nil {
		log.Fatal(err) // better approach than panic()
	}

	fmt.Println("Connected to MongoDB!")

	collection = client.Database(dbName).Collection(collectionName)

	// collection instance is ready to use throughout the application
	fmt.Println("Collection instance is ready")
}

// MongoDB Helpers - file

// insert 1 record
func insertOneMovie(movie model.Netflix) {
	res, err := collection.InsertOne(context.TODO(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Record ", res.InsertedID)
}

// update 1 record
func updateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId) // convert string id to primitive.ObjectID
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id} // bson.D can also be used

	update := bson.M{
		"$set": bson.M{
			"watched": true,
		},
	}

	res, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count: ", res.ModifiedCount)
}

// delete 1 record
func deleteOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId) // convert string id to primitive.ObjectID
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id} // bson.D can also be used

	res, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted count: ", res.DeletedCount)
}

// delete all records
func deleteAllMovies() int64 {
	res, err := collection.DeleteMany(context.TODO(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted count: ", res.DeletedCount)
	return res.DeletedCount
}

// get all records
func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cursor.Next(context.TODO()) {
		var movie bson.M
		if err = cursor.Decode(&movie); err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cursor.Close(context.TODO()) // close the cursor once stream of documents has exhausted

	return movies
}

// Actual controller - file

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all movies endpoint hit!")

	// w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")

	allMovies := getAllMovies()

	err := json.NewEncoder(w).Encode(allMovies)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create movie endpoint hit!")

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		log.Fatal(err)
	}

	insertOneMovie(movie)

	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Mark movie as watched endpoint hit!")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	movieId := mux.Vars(r)["id"]

	updateOneMovie(movieId)

	json.NewEncoder(w).Encode("Movie marked as watched! ID: " + fmt.Sprintf("%s", movieId))
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete a movie endpoint hit!")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	movieId := mux.Vars(r)["id"]

	deleteOneMovie(movieId)

	json.NewEncoder(w).Encode("Movie deleted! ID: " + fmt.Sprintf("%s", movieId))
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete a movie endpoint hit!")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovies()

	json.NewEncoder(w).Encode("Movies deleted! Count: " + fmt.Sprintf("%d", count))
}
