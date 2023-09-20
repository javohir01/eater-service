package models

import (
	"time"
)

type PaymentCard struct {
	ID         string
	EaterID    string
	Number     string
	IsVerifed  bool
	CreatedAt  time.Time
}
