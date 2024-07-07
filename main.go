package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	router := server.SetRouter(db)
	router.Listen(Addr)

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := router.ShutdownWithContext(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	default:
		log.Println("Server exiting")
	}
}
