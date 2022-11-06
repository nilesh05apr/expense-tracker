package configs

import ( 
	"context"
	"log"
	"time"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)



func ConnectDB() *mongo.Client {

	// Set client options
	client, err := mongo.NewClient(options.Client().ApplyURI(LoadEnv()))

	if err != nil {		
		log.Fatal(err)
		fmt.Println(LoadEnv())
	}


	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer cancel()

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

//	fmt.Println("Connected to MongoDB!")

	return client
}

var DB *mongo.Client = ConnectDB()

func GetCollection(collection string) *mongo.Collection {
	return DB.Database("expenseDB").Collection(collection)
}