package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func main()  {
	// Setup client and ensure connection to mongo
	client, err := mongo.NewClient(options.Client())
	if err != nil {
		log.Fatal(fmt.Sprintf("Mongo err - %+v", err))
	}
	mongoContext, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(mongoContext)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error connecting to mongo - %+v", err))
	}
	err = client.Ping(mongoContext, readpref.Primary())
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not ping mongo server after opening connection! - %+v", err))
	}

	// working on collections
	collection := client.Database("hydra").Collection("personnel")
	rand.Seed(time.Now().UnixNano())
	insertID := rand.Intn(10000)
	name := "anon"+strconv.Itoa(rand.Intn(10000))
	result, err := collection.InsertOne(mongoContext, bson.M{"id": insertID, "name": name})
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not insert into mongo - %+v", err))
	}
	id := result.InsertedID
	fmt.Println("Id of inserted anon: ", id)

	// querying the collection returns a cursor (iterator over the results)
	// var searchFilter = bson.D{ {"id", primitive.Regex{Pattern: `.*`, Options: ""}} }
	// cursor, err := collection.Find(mongoContext, searchFilter) // this line doesn't really work
	cursor, err := collection.Find(mongoContext, bson.D{}, options.Find()) // fill the bson.D with search params
	//cursor, err := collection.Distinct(mongoContext, "_id", bson.D{})
	if err != nil {
		log.Fatal(fmt.Sprintf("Errro searching the mongo DB - %+v", err))
	}
	defer func(){ _ = cursor.Close(mongoContext) }()
	for cursor.Next(mongoContext) {
		var result struct {
			Id 		int 	`mongo:"id"`
			Name 	string 	`mongo:"name"`
		}
		// var result bson.M // bson.M is a marshalling bson (bson -> map[string]interface{})
		err = cursor.Decode(&result)
		if err != nil {
			log.Fatal(fmt.Sprintf("Error parsing result from DB - %s", err))
		}
		fmt.Printf("found this shit in the DB: %+v\n", result)
	}
	if err = cursor.Err(); err != nil {
		log.Fatal(fmt.Println("Seems like there was a cursor error", err))
	}

	// can decode things right from the DB - at this point we should have at least one anon in the DB
	var anonMF struct {
		Id 		int 	`mongo:"id"`
		Name 	string 	`mongo:"name"`
	}
	filter := bson.M{"_id": id}
	err = collection.FindOne(mongoContext, filter).Decode(&anonMF)
	if err !=nil {
		log.Fatal(fmt.Println("Seems like there was a cursor error or couldn't find the filtered out MF", err))
	}
	fmt.Printf("Here's the MF - %+v\n", anonMF)
}
