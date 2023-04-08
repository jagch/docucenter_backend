package planentrega

import (
	"fmt"

	"github.com/google/uuid"
)

type IDMaritimo struct {
	value string
}

func NewIDMaritimo(value string) (IDMaritimo, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return IDMaritimo{}, fmt.Errorf("%s: %s", "el ID es inválido", value)
	}

	return IDMaritimo{
		value: v.String(),
	}, nil
}

func (id IDMaritimo) String() string {
	return id.value
}

type IDPuertoEntrega struct {
	value int
}

func NewIDPuertoEntrega(value int) (IDPuertoEntrega, error) {
	if value <= 0 {
		return IDPuertoEntrega{}, fmt.Errorf("%s", "el id del puerto de entrega no puede ser menor o igual a cero")
	}

	return IDPuertoEntrega{
		value: value,
	}, nil
}

func (ipe IDPuertoEntrega) Int() int {
	return ipe.value
}

type NroFlota struct {
	value string
}

func NewNroFlota(value string) (NroFlota, error) {
	msgPatternDescription := "3 letras iniciales, seguidas de 4 números y finalizando con una letra"

	matched, err := validateFieldWithRegexp("nro_flota", "[a-zA-Z]{3}[0-9]{4}[a-zA-Z]{1}", value, msgPatternDescription)
	if err != nil {
		return NroFlota{}, fmt.Errorf("%s", err.Error())
	}

	if !matched {
		return NroFlota{}, fmt.Errorf("el nro. de flota debe cumplir el formato%s", msgPatternDescription)
	}

	if value == "" {
		return NroFlota{}, fmt.Errorf("el nro. de flota no puede ser vacío")
	}

	return NroFlota{
		value: value,
	}, nil
}

func (nf NroFlota) String() string {
	return nf.value
}
