package user

import (
	"context"

	"github.com/NicolasMartino/simplebank/core/user/dto"
	db "github.com/NicolasMartino/simplebank/db/sqlc"
)

type UserRetriever struct {
	store db.Store
}

func NewUserRetriever(store db.Store) *UserRetriever {
	return &UserRetriever{
		store: store,
	}
}

func (rtvr *UserRetriever) RetrieveOneUser(ctx context.Context, ID int64) (userDto dto.UserDto, err error) {
	user, err := rtvr.store.FindUser(ctx, ID)

	if err == nil {
		userDto = dto.UserDto{
			ID:               user.ID,
			Email:            user.Email,
			FirstName:        user.FirstName,
			LastName:         user.LastName,
			PasswordChangeAt: user.PasswordChangeAt,
		}
	}

	return
}
