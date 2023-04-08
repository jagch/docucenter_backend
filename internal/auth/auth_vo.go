package auth

import (
	"errors"
)

var ErrInvalidAuthUsuario = errors.New("usuario inválido")

type AuthUsuario struct {
	value string
}

func NewAuthUsuario(value string) (AuthUsuario, error) {
	if value == "" {
		return AuthUsuario{}, ErrInvalidAuthUsuario
	}

	return AuthUsuario{
		value: value,
	}, nil
}

func (usuario AuthUsuario) String() string {
	return usuario.value
}

var ErrEmptyAuthClave = errors.New("clave no puede ser vacío")

type AuthClave struct {
	value string
}

func NewAuthClave(value string) (AuthClave, error) {
	if value == "" {
		return AuthClave{}, ErrEmptyAuthClave
	}

	return AuthClave{
		value: value,
	}, nil
}

func (clave AuthClave) String() string {
	return clave.value
}

func NewAuth(usuario, clave string) (Auth, error) {
	usuarioVO, err := NewAuthUsuario(usuario)
	if err != nil {
		return Auth{}, err
	}

	claveVO, err := NewAuthClave(clave)
	if err != nil {
		return Auth{}, err
	}

	return Auth{
		Usuario: usuarioVO,
		Clave:   claveVO,
	}, nil
}
