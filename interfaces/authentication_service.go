package interfaces

import "leanmeal/api/dtos"

type AuthenticationService interface {
	Start()
	GetMessage(email string) dtos.InitAuthReponse
	VerifySignature(response dtos.FinishAuthResponse) bool
}
