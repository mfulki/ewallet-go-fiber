package server

import "github.com/gofiber/fiber/v2"

func SetupRouter(h *Handlers) *fiber.App{
	router := fiber.New()

	userRouter :=router.Group("/")
	{
		userRouter.Post("/login", h.AuthHandler.Login)
	}

	return router
}