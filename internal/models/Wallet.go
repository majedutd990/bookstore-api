package models

import "github.com/majedutd990/bookstore-api/internal/models/types"

type (
	Wallet struct {
		Base
		UserId  uint
		Balance types.Price
		Status  types.WalletStatus
	}
)
