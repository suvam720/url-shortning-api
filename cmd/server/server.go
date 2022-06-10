package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/suvam720/url-shortner/pkg/database"
	"github.com/suvam720/url-shortner/pkg/routers"
)

func main() {
	key, err := godotenv.Read(".env")
	if err != nil {
		log.Panic(err)
	}
	dsn := key["DB_STRING"]

	if err := database.Connect(dsn); err != nil {
		log.Fatal(err)
	}
	r := routers.Router()
	r.Run()
}
