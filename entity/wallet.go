package entity

import "github.com/shopspring/decimal"

type Wallet struct {
	Id           int             
	WalletNumber string         
	UserId       int             
	Balance      decimal.Decimal 
}
