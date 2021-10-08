package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type Person struct {
	ID		primitive.ObjectID	`json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname	string			`json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname	string			`json:"lastname,omitempty" bson:"lastname,omitempty"`
}

var client *mongo.Client

func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var person Person
	json.NewDecoder(request.Body).Decode(&person)
	collection := client.Database("wizarddatabase").Collection("people")
	ctx, err := context.WithTimeout(context.Background(), 10*time.Second)
	if err != nil {
		log.Fatal("error writing timeout on: %v", err)
	}
	result, err := collections.InsertOne(ctx, person)
	if err != nil {
		log.Fatal("error on collections.InsertOne: %v", err)
	}
	json.NewEncoder(response).Encode(result)
}

func GetPeopleEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var people []Person
	collection :=  client.Database("wizarddatabase").Collection("poeple")
	ctx, err := context.WithTimeout(context.Background(), 10*time.Second)
	if err != nil {
		log.Fatal("error writing timeout on: %v", err)
	}
	cursor, err := collections.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	if err := defer cursor.Close(ctx); err != nil {
		log.Fatalln("failed to defer the close on ctx: %v", err)
	}
	for cursor.Next(ctx) {
		var person Person
		cursor.Decode(&person)
		people = append(people, person)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(people)
}

func GetPersonEndpoint(response http.ResponseWriter, request *http.Request) {
    response.Header().Add("content-type", "application/json")
    params := mux.Vars(request)
	id := primitive.ObjectIDFromHex(params["id"])
	var person Person
	collection :=  client.Database("wizarddatabase").Collection("people")
    ctx, err := context.WithTimeout(context.Background(), 10*time.Second)
    if err != nil {
		log.Fatal("error writing timeout on: %v", err)
	}
	err := collection.FindOne(ctx, Person{ID: id}).Decode(&person)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
    	response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(person)
}


func main() {
	fmt.Println("Starting the application...")
	ctx, err := context.WithTimeout(context.Background(), 10*time.Second)
	if err != nil {
		log.Fatal("error writing timeout on: %v", err)
	}
	client, err = mongo.Connect(ctx, "mongodb://localhost:27017")
	if err != nil {
		log.Fatal("error connecting to mongodb: %v", err)
	}
	router := mux.NewRouter()
	router.HandleFunc("/person", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/person{id}", GetPersonEndpoint).Methods("GET")
	log.Fatalln(http.ListenAndServe(":12345", router))
}
