package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mfulki/ewallet-go-fiber/constant"
	"github.com/mfulki/ewallet-go-fiber/dto/request"
	"github.com/mfulki/ewallet-go-fiber/dto/response"
	"github.com/mfulki/ewallet-go-fiber/usecase"
	"github.com/mfulki/ewallet-go-fiber/utils"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase: authUsecase,
	}
}

func (h *AuthHandler) Login(ctx *fiber.Ctx) error {
	body := new(request.UserLogin)
	if err := ctx.BodyParser(&body); err != nil {
		return err
	}
	if err := utils.NewValidator().Validate(body); err != nil {
		return err
	}

	token, err := h.authUsecase.Login(body.User(), ctx.Context())
	if err != nil {
		return err

	}

	resp := response.Body{
		Message: constant.LoginPassedMsg,
		Data:    *token,
	}
	ctx.Status(fiber.StatusOK).JSON(resp)
	return nil

}
