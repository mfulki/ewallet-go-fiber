package main

import (
	"log"

	"github.com/mfulki/ewallet-go-fiber/config"
	database "github.com/mfulki/ewallet-go-fiber/db"
	"github.com/mfulki/ewallet-go-fiber/server"
)

const Addr = ":8081"

func main() {
	err := config.EnvInit()
	if err != nil {
		log.Fatal("error get .env")
	}
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("error connecting to DB: %s", err.Error())
	}
	defer db.Close()
	router := server.NewServer(db).SetupServer()
	router.Listen(Addr)

}
