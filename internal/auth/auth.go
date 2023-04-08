package auth

import (
	"context"
)

type AuthUsecase interface {
	Auth(ctx context.Context, auth Auth) (int, error)
}

type AuthStorage interface {
	Get(ctx context.Context, auth Auth) (int, error)
}
