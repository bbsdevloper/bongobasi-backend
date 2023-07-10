package user_handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Prosecutor1x/citizen-connect-frontend/database"
	"github.com/Prosecutor1x/citizen-connect-frontend/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const dbName = "user-list"
const colName = "users"

var collection *mongo.Collection

func init() {

	client, err := database.ConnectToMongoDB()
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready")
}

func createUser(user model.UserData) {
	inserted, err := collection.InsertOne(context.Background(), user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", inserted.InsertedID)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var user model.UserData
	_ = json.NewDecoder(r.Body).Decode(&user)

	createUser(user)
	json.NewEncoder(w).Encode(user)

}

func CheckUserExist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var phone model.Phone
	_ = json.NewDecoder(r.Body).Decode(&phone)
	filter := bson.M{"userphone": phone.Phone}

	var result bson.M

	err := collection.FindOne(context.Background(), filter).Decode(&result)
	fmt.Println("result", result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No documents found")
			// send 500 error
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("No documents found")
		} else {
			json.NewEncoder(w).Encode("Error decoding")
		}
	} else {
		json.NewEncoder(w).Encode(result)
	}

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)
	userId := params["id"]
	if userId == "" {
		response := map[string]interface{}{
			"message": "Please provide correct user id",
		}
		w.WriteHeader(http.StatusNotFound) // Set the HTTP status code
		json.NewEncoder(w).Encode(response)
		return

	}
	id, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.M{"_id": id}
	var userData bson.M
	err := collection.FindOne(context.Background(), filter).Decode(&userData)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(userData)

}

//update user data


func UpdateOneUser(userId string, user model.UserData) {
	id, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": user}
	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		//console log error
		log.Panic(err) 
		}
	fmt.Println("Modified Count: ", result.ModifiedCount)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	fmt.Println("update issue handler")
	var user model.UserData
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// Handle JSON decoding error
		response := map[string]interface{}{
			"message": "Failed to decode request body",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	params := mux.Vars(r)
	userID := params["id"]
	if userID == "" {
		response := map[string]interface{}{
			"message": "Please provide a correct issue ID",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Update the issue with the given ID
	UpdateOneUser(userID, user)
	if err != nil {
		// Handle error in updating the issue
		response := map[string]interface{}{
			"message": "Failed to update the issue",
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := map[string]interface{}{
		"message": "Issue updated successfully",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
