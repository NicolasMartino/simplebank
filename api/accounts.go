package api

import (
	"database/sql"
	"net/http"

	db "github.com/NicolasMartino/simplebank/db/sqlc"

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

type GetAccountDTO struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) GetAccount(ctx *gin.Context) {
	var dto GetAccountDTO

	if err := ctx.ShouldBindUri(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.FindAccount(ctx, dto.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, nil)
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type GetAccountsWithPaginationDTO struct {
	PageNumber int32 `form:"page_number" binding:"required,min=1"`
	PageSize   int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) GetAccountsWithPagination(ctx *gin.Context) {
	var dto GetAccountsWithPaginationDTO

	if err := ctx.ShouldBindQuery(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.FindAccountsWithPaginationParams{
		Offset: (dto.PageNumber - 1) * dto.PageSize,
		Limit:  dto.PageSize,
	}

	account, err := server.store.FindAccountsWithPagination(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type UpdateAccountDTO struct {
	Balance float64 `json:"balance" binding:"min=0"`
}

func (server *Server) UpdateAccount(ctx *gin.Context) {
	var dtoId GetAccountDTO
	var dto UpdateAccountDTO

	if err := ctx.ShouldBindUri(&dtoId); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.UpdateAccountParams{
		ID:      dtoId.ID,
		Balance: dto.Balance,
	}

	account, err := server.store.UpdateAccount(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (server *Server) DeleteAccount(ctx *gin.Context) {
	var dto GetAccountDTO

	if err := ctx.ShouldBindUri(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteAccount(ctx, dto.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
