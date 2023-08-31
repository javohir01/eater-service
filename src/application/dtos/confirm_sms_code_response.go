package dtos

import "github.com/javohir01/eater-service/src/domain/eater/models"

type ComfirmSMSCodeResponse struct {
	Token string `json:"token"`
	Profile
}
