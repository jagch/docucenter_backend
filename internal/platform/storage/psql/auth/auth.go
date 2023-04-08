package auth

import (
	"context"
	"database/sql"
	"fmt"
	"jagch/backend/internal/auth"
)

type AuthStorage struct {
	db *sql.DB
}

func NewAuthStorage(db *sql.DB) *AuthStorage {
	return &AuthStorage{
		db: db,
	}
}

func (r *AuthStorage) Get(ctx context.Context, auth auth.Auth) (int, error) {
	stmt, err := r.db.PrepareContext(ctx, "SELECT u.id FROM public.usuario u WHERE u.usuario = $1 and u.clave = $2")
	if err != nil {
		return 0, fmt.Errorf("error trying get credentials on database: %v", err)
	}

	defer stmt.Close()

	var id int
	usuario := auth.Usuario.String()
	clave := auth.Clave.String()
	if err := stmt.QueryRowContext(ctx, usuario, clave).Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}

		return 0, fmt.Errorf("error trying get credentials on database: %v", err)
	}

	return id, nil
}
