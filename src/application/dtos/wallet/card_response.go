package wallet

import "time"

type CardResponse struct {
	ID         string    `json:"id"`
	EaterID    string    `json:"eater_id"`
	Number     string    `json:"number"`
	IsVerified bool      `json:"is_verified"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewCardResponse(id, eaterID, number string, isVerified bool, createdAt, updatedAt time.Time) *CardResponse {
	return &CardResponse{
		ID:         id,
		EaterID:    eaterID,
		Number:     number,
		IsVerified: isVerified,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
	}
}
