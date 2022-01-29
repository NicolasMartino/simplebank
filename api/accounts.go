package api

import (
	"net/http"
	db "simplt_bank/db/sqlc"

	"github.com/gin-gonic/gin"
)

type CreateAccountDTO struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=EUR USD"`
}

func (server *Server) CreateAccount(ctx *gin.Context) {
	var dto CreateAccountDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner:    dto.Owner,
		Currency: dto.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}
