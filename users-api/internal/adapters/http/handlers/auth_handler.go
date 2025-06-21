package handlers

import (
	"net/http"

	"github.com/PyMarcus/FreelaIF/users-api/internal/adapters/repositories"
	httpport "github.com/PyMarcus/FreelaIF/users-api/internal/core/ports/http"
	"github.com/PyMarcus/FreelaIF/users-api/internal/core/usecases"
	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	usecase *usecases.AuthUsecase
}

func NewLoginHandler(usecase *usecases.AuthUsecase) httpport.AuthHandler {
	return &LoginHandler{
		usecase: usecase,
	}
}

func (lh *LoginHandler) GetToken(ctx *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user, err := lh.usecase.GetUser(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token := lh.createToken(user)

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (lh *LoginHandler) createToken(user repositories.User) string {
	return "fake-jwt-token"
}
