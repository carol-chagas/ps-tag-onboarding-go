package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"ps-tag-onboarding-go/config"
	"ps-tag-onboarding-go/internal/handler"
	"ps-tag-onboarding-go/internal/repository"
	"ps-tag-onboarding-go/internal/service"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient, err := config.GetMongoClient(ctx, "mongodb://mongo:27017")
	if err != nil {
		log.Fatal(err)
	}
	db := mongoClient.Database("tag-onboarding")

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	http.HandleFunc("/save", userHandler.SaveUser)
	http.HandleFunc("/find", userHandler.FindUserByID)
	http.HandleFunc("/users", userHandler.GetUsers)
	http.HandleFunc("/delete", userHandler.DeleteUser)
	http.HandleFunc("/update", userHandler.UpdateUser)

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
