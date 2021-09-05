package types

import (
	"fmt"
	"github.com/majedutd990/bookstore-api/pkg/errors"
	"github.com/majedutd990/bookstore-api/pkg/translate/messages"
)

type (
	BookStatus uint
)

var (
	bookStatus = map[BookStatus]string{
		BookPending:   "pending",
		BookConfirmed: "confirmed",
		BookReject:    "reject",
		BookClose:     "close",
	}
)

const (
	_ BookStatus = iota
	BookPending
	BookConfirmed
	BookReject
	BookClose
)

func (b BookStatus) String() string {
	if s, ok := bookStatus[b]; ok {
		return s
	}
	return fmt.Sprintf("BookStatus(#{b})")
}
func (b BookStatus) MarshalText() ([]byte, error) {
	return []byte(b.String()), nil
}

func (b *BookStatus) UnMarshalText(by []byte) error {
	for i, v := range bookStatus {
		if v == string(by) {
			*b = i
			return nil
		}
	}
	return errors.New(errors.KindInvalid, messages.InvalidBookStatus)
}
