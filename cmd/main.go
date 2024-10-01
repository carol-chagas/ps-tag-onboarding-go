package main

import (
	"log"
	"net/http"

	"ps-tag-onboarding-go/config"
	"ps-tag-onboarding-go/internal/handler"
	"ps-tag-onboarding-go/internal/repository"
	"ps-tag-onboarding-go/internal/service"
)

func main() {
	mongoClient := config.ConnectMongoDB("mongodb://localhost:27017")
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
