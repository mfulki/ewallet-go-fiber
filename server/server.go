package server

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func SetRouter(db *sql.DB) *fiber.App {

	router := fiber.New()

	return router
}
