package cliente

import (
	"context"
)

type ClienteStorage interface {
	Create(ctx context.Context, cliente Cliente) (ClienteResponse, error)
	GetAll(ctx context.Context) (ClientesResponse, error)
}
