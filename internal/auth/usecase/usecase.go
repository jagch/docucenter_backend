package usecase

import (
	"context"
	"database/sql"
	"jagch/backend/internal/auth"
)

type AuthUsecase struct {
	authStorage auth.AuthStorage
}

func NewUsecase(authStorage auth.AuthStorage) AuthUsecase {
	return AuthUsecase{
		authStorage: authStorage,
	}
}

func (a AuthUsecase) Auth(ctx context.Context, auth auth.Auth) (int, error) {
	id, err := a.authStorage.Get(ctx, auth)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}

		return 0, err
	}

	return id, nil
}
