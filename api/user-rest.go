package api

import (
	"net/http"

	customvalidator "github.com/NicolasMartino/simplebank/api/custom-validator"
	"github.com/NicolasMartino/simplebank/core/user/dto"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (server *Server) CreateUser(ctx *gin.Context) {
	var dto dto.CreateUserDTO

	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ok := customvalidator.PasswordValidator(dto.Password, server.appConfig.MinPasswordEntropy)

	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Password entropy is too low"})
		return
	}

	account, err := server.services.UserPersister.CreateUser(ctx, dto)

	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
			}
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (server *Server) FindUser(ctx *gin.Context) {
	var uri URI

	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userDto, err := server.services.UserRetriever.RetrieveOneUser(ctx, uri.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, userDto)
}
