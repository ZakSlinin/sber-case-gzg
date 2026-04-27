package models

type EmailVerificationError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
