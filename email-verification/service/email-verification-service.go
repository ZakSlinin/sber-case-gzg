package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/email-verification/mailer"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

type EmailVerificationService interface {
	VerifyEmail(ctx context.Context, req string) error
	ConfirmEmail(ctx context.Context, tokenString string) error
}

type emailVerificationService struct {
	mailer          mailer.EmailVerificationMailer
	urlVerification string
	notifyURL       string
}

func NewEmailVerificationService(mailer mailer.EmailVerificationMailer, urlVerification string, notifyUrl string) *emailVerificationService {
	return &emailVerificationService{mailer: mailer, urlVerification: urlVerification, notifyURL: notifyUrl}
}

func (s *emailVerificationService) VerifyEmail(ctx context.Context, email string) error {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(s.urlVerification))

	body := fmt.Sprintf(`Привет!
		чтобы поставить отзыв необходимо подтвердить email. 
		Для этого перейдите по ссылке: %s

		Если письмо отправлено по ошибке, проигнорируйте его.
	`, s.urlVerification+tokenString)

	err := s.mailer.Send(email, "Подтвердите email", body)

	if err != nil {
		return err
	}

	return nil
}

func (s *emailVerificationService) ConfirmEmail(ctx context.Context, tokenString string) error {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.urlVerification), nil
	})
	if err != nil || !token.Valid {
		return fmt.Errorf("токен недействителен")
	}

	claims := token.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	return s.notifyOtherService(ctx, email)
}

func (s *emailVerificationService) notifyOtherService(ctx context.Context, email string) error {
	body, _ := json.Marshal(map[string]interface{}{
		"email":   email,
		"correct": true,
	})

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.notifyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("service response%d", resp.StatusCode)
	}

	return nil
}
