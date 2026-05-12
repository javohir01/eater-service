package models

import "time"

type PaymentCard struct {
	ID         string
	EaterID    string
	Number     string
	CardToken  string
	IsVerified bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
