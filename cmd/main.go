package main

import (
	"log"
	"net/http"

	"ps-tag-onboarding-go/internal/handler"
	"ps-tag-onboarding-go/internal/repository"
	"ps-tag-onboarding-go/internal/service"
)

func main() {
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	http.HandleFunc("/save", userHandler.SaveUser)
	http.HandleFunc("/find/", userHandler.FindUserByID)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
