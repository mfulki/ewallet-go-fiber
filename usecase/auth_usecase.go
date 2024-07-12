package usecase

import (
	"context"
	"os"
	"strconv"

	"github.com/mfulki/ewallet-go-fiber/apperrors"
	"github.com/mfulki/ewallet-go-fiber/entity"
	"github.com/mfulki/ewallet-go-fiber/repository"
	"github.com/mfulki/ewallet-go-fiber/utils"
)

type AuthUsecase interface {
	Login(user entity.User, ctx context.Context) (*string, error)
	Register(user *entity.User, ctx context.Context) (*entity.User, error)
	ForgotPassword(ctx context.Context, user *entity.User) (*entity.TokenPassword, error)
	ResetPassword(ctx context.Context, token string, password string) error
}

type authUsecaseImpl struct {
	userRepository          repository.UserRepository
	walletRepository        repository.WalletRepository
	passwordTokenRepository repository.PasswordTokenRepository
}

func NewAuthUsecaseImpl(userRepository repository.UserRepository, walletRepository repository.WalletRepository, passwordTokenRepository repository.PasswordTokenRepository) *authUsecaseImpl {
	return &authUsecaseImpl{
		userRepository:          userRepository,
		walletRepository:        walletRepository,
		passwordTokenRepository: passwordTokenRepository,
	}
}

func (u *authUsecaseImpl) Login(user entity.User, ctx context.Context) (*string, error) {

	password := user.Password
	users, err := u.userRepository.Login(user, ctx)
	if err != nil {
		return nil, err

	}

	hashPassword := []byte(users.Password)
	condition, _ := utils.CheckPassword(password, hashPassword)

	if !condition {
		return nil, apperrors.ErrUnAuthorized()
	}
	token, _ := utils.CreateAccessToken(users, os.Getenv("SECRET_CODE"))
	return &token, nil
}
func (u *authUsecaseImpl) Register(user *entity.User, ctx context.Context) (*entity.User, error) {
	userId, _ := u.userRepository.GetRegisteredUserIdByEmail(user.Email, ctx)
	if userId != nil {
		return nil, apperrors.ErrBadRequest()
	}
	cost, err := strconv.Atoi(os.Getenv("HASH_COST"))
	if err != nil {
		return nil, err
	}

	hash, _ := utils.HashPassword(user.Password, cost)
	stringPassword := string(hash)
	user.Password = stringPassword

	userRegistered, err := u.userRepository.RegisterUser(user, ctx)
	if err != nil {
		return nil, err
	}

	walletDigit := 9900000000000 + user.Id
	walletDigitString := strconv.Itoa(walletDigit)
	walet := entity.Wallet{WalletNumber: walletDigitString, UserId: userRegistered.Id}
	wallets, err := u.walletRepository.RegisterWallet(&walet, ctx)
	user.Wallet = wallets
	if err != nil {
		return nil, err

	}
	return userRegistered, nil
}

func (u *authUsecaseImpl) ForgotPassword(ctx context.Context, user *entity.User) (*entity.TokenPassword, error) {
	user, err := u.userRepository.GetUserIdFromEmail(ctx, user)
	if err != nil {
		return nil, err
	}
	length, err := strconv.Atoi(os.Getenv("RANDOM_LENGTH"))
	if err != nil {
		return nil, err
	}
	token := utils.InitRandomToken(length)
	if err = u.passwordTokenRepository.DeleteToken(ctx, user.Id); err != nil {
		return nil, err
	}
	tokenPassword, err := u.passwordTokenRepository.InsertToken(ctx, user.Id, token)
	if err != nil {
		return nil, err
	}

	return tokenPassword, nil
}
func (u *authUsecaseImpl) ResetPassword(ctx context.Context, token string, password string) error {
	userId, err := u.passwordTokenRepository.CheckToken(ctx, token)
	if err != nil {
		return err
	}
	if err = u.passwordTokenRepository.DeleteToken(ctx, *userId); err != nil {
		return err
	}
	cost, err := strconv.Atoi(os.Getenv("HASH_COST"))
	if err != nil {
		return err
	}
	hash, _ := utils.HashPassword(password, cost)
	stringPassword := string(hash)
	if err = u.userRepository.ResetPassword(ctx, stringPassword, *userId); err != nil {
		return err
	}
	return nil

}
