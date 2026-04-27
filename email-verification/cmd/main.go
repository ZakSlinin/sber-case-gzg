package main

import (
	"github.com/email-verification/handler"
	"github.com/email-verification/mailer"
	"github.com/email-verification/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	godotenv.Load()

	smtpMailer := mailer.NewEmailVerificationMailer(
		os.Getenv("SMTP_HOST"),
		os.Getenv("SMTP_PORT"),
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PASS"),
	)

	//urlVerification := os.Getenv("URL_VERIFICATION")
	urlVerification := "http://localhost:8080/api/email-verification/verify?token="

	emailVerificationService := service.NewEmailVerificationService(smtpMailer, urlVerification)
	emailVerificationHandler := handler.NewEmailVerificationHandler(emailVerificationService)

	r := gin.Default()

	apiEmailVerification := r.Group("/api/email-verification")
	{
		apiEmailVerification.POST("/verify", emailVerificationHandler.VerifyEmail)
		apiEmailVerification.GET("/verify", emailVerificationHandler.ConfirmEmail)
	}

	r.Run(":8080")
}
