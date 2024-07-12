package repository

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/mfulki/ewallet-go-fiber/apperrors"
	"github.com/mfulki/ewallet-go-fiber/db"
	"github.com/mfulki/ewallet-go-fiber/entity"
	"github.com/shopspring/decimal"
)

type WalletRepository interface {
	RegisterWallet(wallet *entity.Wallet, ctx context.Context) (*entity.Wallet, error)
	LockBallance(ctx context.Context, walletId int) (*decimal.Decimal, error)
	AdditionBalance(ctx context.Context, amount decimal.Decimal, walletId int) (*decimal.Decimal, error)
	SubtractionBalance(ctx context.Context, amount decimal.Decimal, walletId int) (*decimal.Decimal, error)
	GetWalletIdByUserId(ctx context.Context, userId int) (*int, error)
	GetWalletIdByWalletNumber(ctx context.Context, walletNumber int) (*int, error)
	GetWalletDetails(ctx context.Context, userId int) (*entity.Wallet, error)
}

type walletRepositoryImpl struct {
	db db.DBConnection
}

func NewWalletRepository(db db.DBConnection) *walletRepositoryImpl {
	return &walletRepositoryImpl{
		db: db,
	}
}

func (r *walletRepositoryImpl) RegisterWallet(wallet *entity.Wallet, ctx context.Context) (*entity.Wallet, error) {
	q := `insert into wallets (wallet_number,user_id)
		  values ($1,$2) returning id;`
	err := r.db.QueryRowContext(ctx, q, wallet.WalletNumber, wallet.UserId).Scan(&wallet.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.ErrNotFound()
		}
		return nil, apperrors.ErrInternalServer()
	}
	return wallet, err
}
func (r *walletRepositoryImpl) LockBallance(ctx context.Context, walletId int) (*decimal.Decimal, error) {
	tx := ctx.Value("db-tx").(*sql.Tx)
	var balance *decimal.Decimal
	q := `select balance from wallets where id=$1 for update;`
	err := tx.QueryRowContext(ctx, q, walletId).Scan(&balance)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.ErrNotFound()
		}
		return nil, apperrors.ErrInternalServer()
	}
	return balance, err
}

func (r *walletRepositoryImpl) AdditionBalance(ctx context.Context, amount decimal.Decimal, walletId int) (*decimal.Decimal, error) {
	tx := ctx.Value("db-tx").(*sql.Tx)
	var balance *decimal.Decimal
	q := `update wallets set balance=balance+$1 where id =$2 returning balance;`
	err := tx.QueryRowContext(ctx, q, amount, walletId).Scan(&balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.ErrNotFound()
		}
		return nil, apperrors.ErrInternalServer()
	}
	return balance, err
}
func (r *walletRepositoryImpl) SubtractionBalance(ctx context.Context, amount decimal.Decimal, walletId int) (*decimal.Decimal, error) {
	tx := ctx.Value("db-tx").(*sql.Tx)
	var balance *decimal.Decimal
	q := `update wallets set balance=balance-$1 where id =$2 returning balance;`
	err := tx.QueryRowContext(ctx, q, amount, walletId).Scan(&balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.ErrNotFound()
		}
		return nil, apperrors.ErrInternalServer()
	}
	return balance, err
}
func (r *walletRepositoryImpl) GetWalletIdByUserId(ctx context.Context, userId int) (*int, error) {
	q := `select id from wallets where user_id=$1;`
	var walletId *int
	err := r.db.QueryRowContext(ctx, q, userId).Scan(&walletId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.ErrNotFound()
		}
		return nil, apperrors.ErrInternalServer()
	}
	return walletId, err
}
func (r *walletRepositoryImpl) GetWalletIdByWalletNumber(ctx context.Context, walletNumber int) (*int, error) {
	q := `select id from wallets where wallet_number=$1;`
	var walletId *int
	walletnumberstr := strconv.Itoa(walletNumber)
	err := r.db.QueryRowContext(ctx, q, walletnumberstr).Scan(&walletId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.ErrNotFound()
		}
		return nil, apperrors.ErrInternalServer()
	}
	return walletId, err
}
func (r *walletRepositoryImpl) GetWalletDetails(ctx context.Context, userId int) (*entity.Wallet, error) {
	var wallet entity.Wallet
	q := `select wallet_number,balance from wallets where user_id=$1`
	err := r.db.QueryRowContext(ctx, q, userId).Scan(&wallet.WalletNumber, &wallet.Balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.ErrNotFound()
		}
		return nil, apperrors.ErrInternalServer()
	}
	return &wallet, err
}
