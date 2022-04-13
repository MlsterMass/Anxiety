package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MlsterMass/Anxiety/repository"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	userCollection       *mongo.Collection
	volunteerCollection  *mongo.Collection
	specialistCollection *mongo.Collection
	supportCollection    *mongo.Collection
	authErr              = map[string]string{"message": "User doesn't exists"}
	authOK               = map[string]string{"message": "Auth successfully"}
	userOK               = map[string]string{"message": "user created"}
	volunteerOK          = map[string]string{"message": "volunteer created"}
	supportOK            = map[string]string{"message": "support created"}
	specialistOK         = map[string]string{"message": "specialist created"}
)

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
	credential := options.Credential{
		Username: os.Getenv("DB_ADMIN"),
		Password: os.Getenv("DB_ADMIN_PASSWORD"),
	}
	clientOptions :=
		options.Client().ApplyURI(connectionString).SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to mongodb!")

	userCollection = client.Database(dbName).Collection("users")
	volunteerCollection = client.Database(dbName).Collection("volunteers")
	specialistCollection = client.Database(dbName).Collection("specialist")
	supportCollection = client.Database(dbName).Collection("support")

	fmt.Println("collection instance created")
}

func RegistrationUser(w http.ResponseWriter, r *http.Request) {
	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		fmt.Fprintf(w, "error: %v\n", errors.New("reading error"))
	}
	defer r.Body.Close()

	user := repository.Users{}
	jsonErr := json.Unmarshal(body, &user)
	if jsonErr != nil {
		fmt.Fprintf(w, "error: %v\n", errors.New("not supported format"))
	}
	user = repository.Users{
		ID:       primitive.NewObjectID(),
		Name:     user.Name,
		Nickname: user.Nickname,
		Password: user.Password,
		TgAcc:    user.TgAcc,
		Gender:   user.Gender,
		Status:   user.Status,
		Children: user.Children,
		Pets:     user.Pets,
		Location: user.Location,
	}
	err := repository.ValidateUserStruct(user)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	userCreated, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error : %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userOK)
	fmt.Printf("userCreated : %v\n", userCreated.InsertedID)
}

func RegistrationVolunteer(w http.ResponseWriter, r *http.Request) {
	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		fmt.Fprintf(w, "error: %v\n", errors.New("reading error"))
	}
	defer r.Body.Close()

	volunteer := repository.Volunteers{}
	jsonErr := json.Unmarshal(body, &volunteer)
	if jsonErr != nil {
		fmt.Fprintf(w, "error: %v\n", errors.New("not supported format"))
	}
	volunteer = repository.Volunteers{
		ID:       primitive.NewObjectID(),
		Name:     volunteer.Name,
		Nickname: volunteer.Nickname,
		Password: volunteer.Password,
		Gender:   volunteer.Gender,
		Location: volunteer.Location,
	}
	err := repository.ValidateVolunteerStruct(volunteer)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	volunteerCreated, err := volunteerCollection.InsertOne(context.Background(), volunteer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error : %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(volunteerOK)
	fmt.Printf("volunteerCreated : %v\n", volunteerCreated.InsertedID)
}
func RegistrationSpecialist(w http.ResponseWriter, r *http.Request) {
	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		fmt.Fprintf(w, "error: %v\n", errors.New("reading error"))
	}
	defer r.Body.Close()

	specialist := repository.Specialists{}
	jsonErr := json.Unmarshal(body, &specialist)
	if jsonErr != nil {
		fmt.Fprintf(w, "error: %v\n", errors.New("not supported format"))
	}
	specialist = repository.Specialists{
		ID:        primitive.NewObjectID(),
		Firstname: specialist.Firstname,
		Lastname:  specialist.Lastname,
		Nickname:  specialist.Nickname,
		Gender:    specialist.Gender,
		Location:  specialist.Location,
		Sphere:    specialist.Sphere,
		Password:  specialist.Password,
	}
	err := repository.ValidateSpecialistStruct(specialist)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	specialistCreated, err := specialistCollection.InsertOne(context.Background(), specialist)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error : %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(specialistOK)
	fmt.Printf("specialistCreated : %v\n", specialistCreated.InsertedID)
}
func RegistrationSupport(w http.ResponseWriter, r *http.Request) {
	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		fmt.Fprintf(w, "error: %v\n", errors.New("reading error"))
	}
	defer r.Body.Close()

	support := repository.Support{}
	jsonErr := json.Unmarshal(body, &support)
	if jsonErr != nil {
		fmt.Fprintf(w, "error: %v\n", errors.New("not supported format"))
	}
	support = repository.Support{
		ID:       primitive.NewObjectID(),
		Name:     support.Name,
		Nickname: support.Nickname,
		Password: support.Password,
		Gender:   support.Gender,
		Location: support.Location,
	}
	err := repository.ValidateSupportStruct(support)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	supportCreated, err := supportCollection.InsertOne(context.Background(), support)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error : %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(supportOK)
	fmt.Printf("supportCreated : %v\n", supportCreated.InsertedID)
}
func LoginUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "error: %v\n", errors.New("reading error"))
	}
	defer r.Body.Close()
	user := repository.Users{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Fprintf(w, "error: %v\n", errors.New("not supported format"))
	}
	user = repository.Users{
		Nickname: user.Nickname,
		Password: user.Password,
	}
	err = userCollection.FindOne(context.TODO(),
		bson.D{{"nickname", user.Nickname},
			{"password", user.Password},
		}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(authErr)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(authOK)
	fmt.Print(authOK["message"])
}
func LoginVolunteer(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "error: %v\n", errors.New("reading error"))
	}
	defer r.Body.Close()
	volunteer := repository.Volunteers{}
	err = json.Unmarshal(body, &volunteer)
	if err != nil {
		fmt.Fprintf(w, "error: %v\n", errors.New("not supported format"))
	}
	volunteer = repository.Volunteers{
		Nickname: volunteer.Nickname,
		Password: volunteer.Password,
	}
	err = volunteerCollection.FindOne(context.TODO(),
		bson.D{{"nickname", volunteer.Nickname},
			{"password", volunteer.Password},
		}).Decode(&volunteer)
	if err == mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(authErr)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(authOK)
	fmt.Print(authOK["message"])
}
func LoginSpecialist(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "error: %v\n", errors.New("reading error"))
	}
	defer r.Body.Close()
	specialist := repository.Specialists{}
	err = json.Unmarshal(body, &specialist)
	if err != nil {
		fmt.Fprintf(w, "error: %v\n", errors.New("not supported format"))
	}
	specialist = repository.Specialists{
		Nickname: specialist.Nickname,
		Password: specialist.Password,
	}
	err = specialistCollection.FindOne(context.TODO(),
		bson.D{{"nickname", specialist.Nickname},
			{"password", specialist.Password},
		}).Decode(&specialist)
	if err == mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(authErr)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(authOK)
	fmt.Print(authOK["message"])
}

func LoginSupport(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "error: %v\n", errors.New("reading error"))
	}
	defer r.Body.Close()
	support := repository.Support{}
	err = json.Unmarshal(body, &support)
	if err != nil {
		fmt.Fprintf(w, "error: %v\n", errors.New("not supported format"))
	}
	support = repository.Support{
		Nickname: support.Nickname,
		Password: support.Password,
	}
	err = supportCollection.FindOne(context.TODO(),
		bson.D{{"nickname", support.Nickname},
			{"password", support.Password},
		}).Decode(&support)
	if err == mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(authErr)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(authOK)
	fmt.Print(authOK["message"])
}

func Logout(w http.ResponseWriter, r *http.Request) {

}
