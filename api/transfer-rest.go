package api

import (
	"net/http"

	"github.com/NicolasMartino/simplebank/core/transfer"
	"github.com/gin-gonic/gin"
)

func (server *Server) createTransfer(ctx *gin.Context) {
	var dto transfer.TransferDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	transfer, err := server.services.TransferPersister.CreateTransfer(ctx, dto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, transfer)
}
