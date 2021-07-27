package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"example.com/APIDemo/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email    string             `json:"Email,omitempty" bson:"Email,omitempty"`
	Name     string             `json:"Name,omitempty" bson:"Name,omitempty"`
	Type     string             `json:"Type,omitempty" bson:"Type,omitempty"`
	Password []byte             `json:"Password,omitempty" bson:"Password,omitempty"`
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API Works!")
	fmt.Println("Endpoint Hit: homePage")
}

// Register new user in database
//
func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: register")

	w.Header().Set("content-type", "application/json")

	var u User
	json.NewDecoder(r.Body).Decode(&u)

	password, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 14)

	u.Password = password

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := database.GetCollection().InsertOne(ctx, u)
	if err != nil {
		log.Fatal(err.Error())
	}

	json.NewEncoder(w).Encode(result)
}

// Check for user in database, then compare password
//
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: login")

	w.Header().Set("content-type", "application/json")

	var u User
	json.NewDecoder(r.Body).Decode(&u)
	email := u.Email
	password := u.Password

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := database.GetCollection().FindOne(ctx, User{Email: email}).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{ "message": "user not found" }`))
		return
	}

	err = bcrypt.CompareHashAndPassword(u.Password, password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{ "message": "incorrect password" }`))
		return
	}

	json.NewEncoder(w).Encode(u)
}
