package wallet

type AddCardRequest struct {
	Number    string `json:"number" binding:"required"`
	CardToken string `json:"card_token" binding:"required"`
}
