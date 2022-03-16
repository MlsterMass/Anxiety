package handlers

import (
	"Anxiety/repository"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"strconv"
)

var collection *mongo.Collection

func init() {
	loadTheEnv()
	createDBInstance()
}

func loadTheEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading the .env file")
	}
}

func createDBInstance() {
	connectionString := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	collName := os.Getenv("DB_USERS")

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to mongodb!")

	collection = client.Database(dbName).Collection(collName)
	fmt.Println("collection instance created")
}

func RegistrationUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	children, err := strconv.ParseBool(r.FormValue("children"))
	if err != nil {
		fmt.Println("Incorrect value")
	}
	pets, err := strconv.ParseBool(r.FormValue("pets"))
	if err != nil {
		fmt.Println("Incorrect value")
	}
	user := repository.Users{
		ID:       primitive.NewObjectID(),
		Name:     r.FormValue("name"),
		Nickname: r.FormValue("nickname"),
		Gender:   r.FormValue("gender"),
		Status:   r.FormValue("status"),
		Children: children,
		Pets:     pets,
		Location: r.FormValue("location"),
		Password: r.FormValue("password"),
	}
	w.WriteHeader(http.StatusOK)
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("User added")

}

func Auth(w http.ResponseWriter, r *http.Request) {
	user := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Specialist: %v\n", user)

}

func Logout(w http.ResponseWriter, r *http.Request) {
	user := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Specialist: %v\n", user)

}
