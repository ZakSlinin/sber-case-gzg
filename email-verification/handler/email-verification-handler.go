package handler

import (
	"github.com/email-verification/models"
	"github.com/email-verification/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EmailVerificationHandler interface {
	VerifyEmail(emailVerificationService service.EmailVerificationService) error
	ConfirmEmail(g *gin.Context)
}

type emailVerificationHandler struct {
	emailVerificationService service.EmailVerificationService
}

func NewEmailVerificationHandler(emailVerificationService service.EmailVerificationService) *emailVerificationHandler {
	return &emailVerificationHandler{emailVerificationService: emailVerificationService}
}

func (h *emailVerificationHandler) VerifyEmail(g *gin.Context) {
	var req models.EmailVerificationRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, &models.EmailVerificationError{Code: 400, Message: err.Error()})
		return
	}

	if err := h.emailVerificationService.VerifyEmail(g.Request.Context(), req.Email); err != nil {
		g.JSON(http.StatusBadRequest, &models.EmailVerificationError{Code: 400, Message: err.Error()})
		return
	}

	g.JSON(http.StatusOK, "correct")
}

func (h *emailVerificationHandler) ConfirmEmail(g *gin.Context) {
	token := g.Query("token")

	if err := h.emailVerificationService.ConfirmEmail(g.Request.Context(), token); err != nil {
		g.JSON(http.StatusBadRequest, &models.EmailVerificationError{Code: 400, Message: err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "Email подтверждён"})
}
