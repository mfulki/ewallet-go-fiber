package request

import "github.com/mfulki/ewallet-go-fiber/entity"

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (req *UserLogin) User() entity.User {
	return entity.User{
		Email:    req.Email,
		Password: req.Password,
	}
}
