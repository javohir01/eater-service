packege dtos

typeEaterSignupResponse struct {
	EaterID string `json: "eater_id"`
}

func NewEaterSignupResponse (eaterID string ) * EaterSignupResponse {
	return&EaterSignupResponse{
		EaterID: eaterID
	}
}