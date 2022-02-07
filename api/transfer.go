package api

import (
	"net/http"

	db "github.com/NicolasMartino/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type TransferDTO struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,oneof=USD EUR CAD"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var dto TransferDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.CreateTransferParams{
		FromAccountID: dto.FromAccountID,
		ToAccountID:   dto.ToAccountID,
		Amount:        float64(dto.Amount),
	}

	transfer, err := server.store.CreateTransfer(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, transfer)
}
