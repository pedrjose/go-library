package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var OwnersCollection *mongo.Collection
var BooksCollection *mongo.Collection

func InitDB() error {
	fmt.Println("STARTING GO SERVER...")

	// Cria uma instância para a versão da API do servidor
	fmt.Println("1. Creating a instance for server API")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	// Define uma conexão com o cluster do MongoDB Atlas
	fmt.Println("2. Define a connection to the MongoDB Atlas cluster")
	opts := options.Client().ApplyURI("mongodb+srv://user:user@cluster0.mga5zza.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0").SetServerAPIOptions(serverAPI)

	// Estabelece a conexão com o MongoDB
	fmt.Println("3. Get the MongoDB connection")
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	// Verifica se a conexão está funcionando corretamente
	fmt.Println("4. Checking if the connection is working properly")
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}

	// Define as coleções MongoDB
	fmt.Println("5. Setting Database Collections")
	// Define o modelo de índice para o campo "email" na coleção "owners"
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1},              // Campo de índice e sua ordem (1 para ascendente, -1 para descendente)
		Options: options.Index().SetUnique(true), // Define o índice como único
	}

	// Cria o índice na coleção OwnersCollection
	_, err = client.Database("quickstart").Collection("owners").Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		panic(err)
	}

	// Define as coleções MongoDB
	OwnersCollection = client.Database("quickstart").Collection("owners")
	BooksCollection = client.Database("quickstart").Collection("books")

	fmt.Println("PINGED YOUR DEPLOYMENT: MongoDB connection initialized. You successfully connected!")
	return nil
}

func init() {
	err := InitDB()
	if err != nil {
		panic(err)
	}
}
