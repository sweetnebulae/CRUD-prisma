package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go_prisma/helper"
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

	server := http.Server{
		Addr:           ":" + os.Getenv("PORT"+"\n"),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	serverErr := server.ListenAndServe()
	if serverErr != nil {
		panic(serverErr)
	}
}
