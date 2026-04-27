package service

import (
	"context"
	"fmt"
	"github.com/email-verification/mailer"
	"github.com/google/uuid"
)

type EmailVerificationService interface {
	VerifyEmail(ctx context.Context) error
}

type emailVerificationService struct {
	mailer          mailer.EmailVerificationMailer
	urlVerification string
}

func (s *emailVerificationService) VerifyEmail(ctx context.Context, email string) error {
	verifyToken := uuid.NewString()

	body := fmt.Sprintf(`Привет!
чтобы поставить отзыв необходимо подтвердить email. 
Для этого перейдите по ссылке: %s

Если письмо отправлено по ошибке, проигнорируйте его.
`, s.urlVerification+verifyToken)

	err := s.mailer.Send(email, "Подтвердите email", body)

	if err != nil {
		return err
	}

	return nil
}
