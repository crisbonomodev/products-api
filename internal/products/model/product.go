package model

import (
	"github.com/google/uuid"
)

type Product struct {
	UID uuid.UUID `db:"uid" json:"uid"`
}
