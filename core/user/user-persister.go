package user

import (
	"context"

	"github.com/NicolasMartino/simplebank/core/user/dto"
	db "github.com/NicolasMartino/simplebank/db/sqlc"
	"github.com/NicolasMartino/simplebank/util"
)

type UserPersister struct {
	store db.Store
}

func NewUserPersister(store db.Store) *UserPersister {
	return &UserPersister{
		store,
	}
}

func (prstr *UserPersister) CreateUser(ctx context.Context, createUserDTO dto.CreateUserDTO) (userDto dto.UserDto, err error) {
	hashedPassword, err := util.HashPassword(createUserDTO.Password)
	if err != nil {
		return
	}

	args := db.CreateUserParams{
		Email:          createUserDTO.Email,
		HashedPassword: hashedPassword,
		FirstName:      createUserDTO.FirstName,
		LastName:       createUserDTO.LastName,
	}

	persistedUser, err := prstr.store.CreateUser(ctx, args)

	if err == nil {
		userDto = dto.UserDto{
			ID:               persistedUser.ID,
			Email:            persistedUser.Email,
			FirstName:        persistedUser.FirstName,
			LastName:         persistedUser.LastName,
			PasswordChangeAt: persistedUser.PasswordChangeAt,
		}
	}
	return
}
