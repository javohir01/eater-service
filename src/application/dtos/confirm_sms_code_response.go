package dtos

import "github.com/javohir01/eater-service/src/domain/eater/models"

type ConfirmSmsCodeResponse struct {
	Token   string               `json:"token"`
	Profile *models.EaterProfile `json:"profile"`
}

type NewConfirmSmsCodeResponse struct {
	Token   string               `json:"token"`
	Profile *models.EaterProfile `json:"profile"`
}
