package api

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kumareddy14322/fullstack/api/controllers"
	"github.com/kumareddy14322/fullstack/api/seed"
)

var server = controllers.Server{}

func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	
	//server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	
	server.Initialize(os.Getenv("postgres"), os.Getenv("postgres"), os.Getenv("kumar"), os.Getenv("5432"), os.Getenv("127.0.0.1"), os.Getenv("fullstack_api"))
	
	seed.Load(server.DB)

	server.Run(":8080")

}