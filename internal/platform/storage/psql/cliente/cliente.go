package cliente

import (
	"context"
	"database/sql"
	"fmt"
	"jagch/backend/internal/cliente"
)

type ClienteStorage struct {
	db *sql.DB
}

func NewClienteStorage(db *sql.DB) *ClienteStorage {
	return &ClienteStorage{
		db: db,
	}
}

func (r *ClienteStorage) Create(ctx context.Context, c cliente.Cliente) (cliente.ClienteResponse, error) {
	stmt, err := r.db.PrepareContext(ctx, "INSERT INTO cliente(id, nombre) VALUES($1, $2) RETURNING id, nombre")
	if err != nil {
		return cliente.ClienteResponse{}, fmt.Errorf("error al intentar guardar el registro en la base de datos: %v", err)
	}

	defer stmt.Close()

	var id, nombre string
	err = stmt.QueryRowContext(ctx, c.ID.String(), c.Nombre.String()).Scan(&id, &nombre)
	if err != nil {
		return cliente.ClienteResponse{}, fmt.Errorf("error al intentar guardar el registro en la base de datos: %v", err)
	}

	return cliente.ClienteResponse{
		ID:     id,
		Nombre: nombre,
	}, nil
}

func (r *ClienteStorage) GetAll(ctx context.Context) (cliente.ClientesResponse, error) {
	stmt, err := r.db.PrepareContext(ctx, "SELECT id, nombre FROM cliente")
	if err != nil {
		return nil, fmt.Errorf("error al intentar obtener los clientes de la base de datos: %v", err)
	}

	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("error al intentar obtener los clientes de la base de datos: %v", err)
	}

	defer rows.Close()

	var clientesResponse cliente.ClientesResponse
	for rows.Next() {
		var cliente cliente.ClienteResponse
		if err = rows.Scan(&cliente.ID, &cliente.Nombre); err != nil {
			return nil, fmt.Errorf("error al intentar obtener los clientes de la base de datos: %v", err)
		}
		clientesResponse = append(clientesResponse, cliente)
	}

	return clientesResponse, nil
}
