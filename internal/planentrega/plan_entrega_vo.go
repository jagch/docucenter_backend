package planentrega

import (
	"fmt"
	"regexp"
	"time"

	"github.com/google/uuid"
)

type IDCliente struct {
	value string
}

func NewIDCliente(value string) (IDCliente, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return IDCliente{}, fmt.Errorf("%s: %w", "el ID del cliente es inválido", err)
	}

	return IDCliente{
		value: v.String(),
	}, nil
}

func (id IDCliente) String() string {
	return id.value
}

type TipoProducto struct {
	value string
}

func NewTipoProducto(value string) (TipoProducto, error) {
	if value == "" {
		return TipoProducto{}, fmt.Errorf("%s", "el tipo de producto no puede ser vacío")
	}

	return TipoProducto{
		value: value,
	}, nil
}

func (tp TipoProducto) String() string {
	return tp.value
}

type CantidadProducto struct {
	value int
}

func NewCantidadProducto(value int) (CantidadProducto, error) {
	if value <= 0 {
		return CantidadProducto{}, fmt.Errorf("%s", "la cantidad de producto no puede ser menor o igual a cero")
	}

	return CantidadProducto{
		value: value,
	}, nil
}

func (cp CantidadProducto) Int() int {
	return cp.value
}

type FechaRegistro struct {
	value time.Time
}

func NewFechaRegistro(value time.Time) (FechaRegistro, error) {
	return FechaRegistro{
		value: value,
	}, nil
}

func (fr FechaRegistro) Date() time.Time {
	return fr.value
}

type FechaEntrega struct {
	value time.Time
}

func NewFechaEntrega(value time.Time) (FechaEntrega, error) {
	return FechaEntrega{
		value: value,
	}, nil
}

func (fr FechaEntrega) Date() time.Time {
	return fr.value
}

type PrecioEnvio struct {
	value float64
}

func NewPrecioEnvio(value float64) (PrecioEnvio, error) {
	if value <= 0 {
		return PrecioEnvio{}, fmt.Errorf("%s", "el precio de envio no puede ser menor o igual a cero")
	}

	return PrecioEnvio{
		value: value,
	}, nil
}

func (pe PrecioEnvio) Float64() float64 {
	return pe.value
}

type NroGuia struct {
	value string
}

func NewNroGuia(value string) (NroGuia, error) {
	msgPatternDescription := "numero unico alfanumerico de 10 dígitos"

	matched, err := validateFieldWithRegexp("nro_guia", "[a-zA-Z0-9]{10}", value, msgPatternDescription)
	if err != nil {
		return NroGuia{}, fmt.Errorf("%s", err.Error())
	}

	if !matched {
		return NroGuia{}, fmt.Errorf("el nro de guia debe cumplir el formato%s", msgPatternDescription)
	}

	if value == "" {
		return NroGuia{}, fmt.Errorf("%s", "el nro. de guía no puede ser vacío")
	}

	return NroGuia{
		value: value,
	}, nil
}

func (ng NroGuia) String() string {
	return ng.value
}

type Dscto struct {
	value float64
}

func NewDscto(value float64) (Dscto, error) {
	if value < 0 {
		return Dscto{}, fmt.Errorf("el dscto. no puede ser menor a cero")
	}

	return Dscto{
		value: value,
	}, nil
}

func (d Dscto) Float64() float64 {
	return d.value
}

func validateFieldWithRegexp(nameField, pattern, s, descriptionOfPattern string) (bool, error) {
	matched, err := regexp.MatchString(`^`+pattern+`$`, s)
	if err != nil {
		return false, fmt.Errorf("error al validar el campo: %s : %s", nameField, err.Error())
	}

	if !matched {
		return false, fmt.Errorf("error al validar el campo: %s : %s", nameField, descriptionOfPattern)
	}

	return true, nil
}
