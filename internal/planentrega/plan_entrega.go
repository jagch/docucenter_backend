package planentrega

import (
	"context"
)

type PlanEntregaStorage interface {
	Create(ctx context.Context, planEntrega any) (string, error)
	GetAll(ctx context.Context) (any, error)
	Update(ctx context.Context, planEntrega any) (string, error)
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, searchParams map[string]SearchParams, page int) (any, error)
}
