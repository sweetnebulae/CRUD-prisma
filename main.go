package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go_prisma/config"
	"go_prisma/controller"
	"go_prisma/helper"
	"go_prisma/repository"
	"go_prisma/router"
	"go_prisma/service"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	fmt.Println("Start server" + os.Getenv("PORT"+"\n"))

	// Handle DB connection
	db, err := config.ConnectDB()
	helper.ErrorPanic(err)

	defer db.Prisma.Disconnect()

	// repository
	postRepository := repository.NewPostRepositoryImpl(db)

	// service
	postService := service.NewPostServiceImpl(postRepository)

	// controller
	postController := controller.NewPostController(postService)

	// router
	routes := router.NewRouter(postController)

	// server
	server := http.Server{
		Addr:           ":" + os.Getenv("PORT"),
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	serverErr := server.ListenAndServe()
	if serverErr != nil {
		panic(serverErr)
	}
}
