package models

import "github.com/majedutd990/bookstore-api/internal/models/types"

type (
	Book struct {
		Base
		Name          string
		Description   string
		File          string
		SellerID      uint
		CategoryID    uint
		Comments      []Comment
		DownloadCount uint
		Pictures      []Picture
		Status        types.BookStatus
	}
)
