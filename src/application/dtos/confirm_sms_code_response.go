package dtos

import "github.com/javohir01/eater-service/src/domain/eater/models"

type ComfirmSMSCodeResponse struct {
	Token   string               `json:"token"`
	Profile *models.EaterProfile `json:"profile"`
}

func NewComfirmSMSCodeResponse(token string, profile *models.EaterProfile) *ComfirmSMSCodeResponse {
	return &ComfirmSMSCodeResponse{
		Profile: profile,
		Token:   token,
	}
}
