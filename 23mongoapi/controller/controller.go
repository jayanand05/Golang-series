package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"hello/23mongoapi/model"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://jay01541:Jay01541@cluster0.kn9hrf8.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colname = "watchlist"

var collection *mongo.Collection

//connection with MongoDB
func init() {
	
	clientOption := options.Client().ApplyURI(connectionString)

	client, err:= mongo.Connect(context.TODO(), clientOption)
	if err!= nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB is connected")

	collection = client.Database(dbName).Collection(colname)
	fmt.Println("collection instance is ready")
}

//MongoDB helpers
func insertOneMovie(movie model.Netflix){
	insert, err := collection.InsertOne(context.Background(), movie)
      
	if err != nil {
		fmt.Println("Something went wrong")
	}
	fmt.Println("The inserted movie with id", insert.InsertedID)
}

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
    //ObjectIDFromHex creates a new ObjectID from a hex string. 
    //It returns an error if the hex string is not a valid ObjectID.
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err!= nil {
		panic(err)
	}
	fmt.Println("Modified count is:", result.ModifiedCount)
	// result.ModifiedCount is a field that represents the number of documents modified or 
	// updated as a result of an operation performed on a MongoDB collection. 
	// It is typically used when executing an update operation using the MongoDB Go driver.
}
func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id":id}

	deleteCount, err:= collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Delete count is", deleteCount.DeletedCount)
}

func deleteAllMovie() {
	deleteResult, err:= collection.DeleteMany(context.Background(), bson.M{}, nil)
	if err != nil {
		log.Fatal(err)
	} 
	

	fmt.Println("Deleted count is", deleteResult.DeletedCount) 
}

func getAllMovies() []primitive.M{
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err!= nil {
		log.Fatal(err)
	}

	var movies []primitive.M
	for cursor.Next(context.Background()){
		var movie bson.M
		err := cursor.Decode(&movie)
		if err!= nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
   return  movies
}

//controller
func GetMyAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}