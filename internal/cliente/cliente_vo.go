package cliente

import (
	"fmt"

	"github.com/google/uuid"
)

type ClienteID struct {
	value string
}

func NewClienteID(value string) (ClienteID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return ClienteID{}, fmt.Errorf("%s: %s", "el ID del cliente es inválido", value)
	}

	return ClienteID{
		value: v.String(),
	}, nil
}

func (id ClienteID) String() string {
	return id.value
}

type ClienteNombre struct {
	value string
}

func NewClienteNombre(value string) (ClienteNombre, error) {
	if value == "" {
		return ClienteNombre{}, fmt.Errorf("%s: %s", "el nombre del cliente es inválido", value)
	}

	return ClienteNombre{
		value: value,
	}, nil
}

func (name ClienteNombre) String() string {
	return name.value
}

func NewCliente(id, nombre string) (Cliente, error) {
	idVO, err := NewClienteID(id)
	if err != nil {
		return Cliente{}, err
	}

	nameVO, err := NewClienteNombre(nombre)
	if err != nil {
		return Cliente{}, err
	}

	return Cliente{
		ID:     idVO,
		Nombre: nameVO,
	}, nil
}
