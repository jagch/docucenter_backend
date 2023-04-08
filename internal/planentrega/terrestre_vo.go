package planentrega

import (
	"fmt"

	"github.com/google/uuid"
)

type IDTerrestre struct {
	value string
}

func NewIDTerrestre(value string) (IDTerrestre, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return IDTerrestre{}, fmt.Errorf("%s: %s", "el ID es inválido", value)
	}

	return IDTerrestre{
		value: v.String(),
	}, nil
}

func (id IDTerrestre) String() string {
	return id.value
}

type IDBodegaEntrega struct {
	value int
}

func NewIDBodegaEntrega(value int) (IDBodegaEntrega, error) {
	if value <= 0 {
		return IDBodegaEntrega{}, fmt.Errorf("%s", "el id de la bodega no puede ser menor o igual a cero")
	}

	return IDBodegaEntrega{
		value: value,
	}, nil
}

func (ibe IDBodegaEntrega) Int() int {
	return ibe.value
}

type PlacaVehiculo struct {
	value string
}

func NewPlacaVehiculo(value string) (PlacaVehiculo, error) {

	msgPatternDescription := "3 letras iniciales y 3 números finales"

	matched, err := validateFieldWithRegexp("placa_vehiculo", "[a-zA-Z]{3}[0-9]{3}", value, msgPatternDescription)
	if err != nil {
		return PlacaVehiculo{}, fmt.Errorf("%s", err.Error())
	}

	if !matched {
		return PlacaVehiculo{}, fmt.Errorf("la placa del vehiculo debe cumplir el formato%s", msgPatternDescription)
	}

	if value == "" {
		return PlacaVehiculo{}, fmt.Errorf("la placa del vehiculo no puede ser vacía")
	}

	return PlacaVehiculo{
		value: value,
	}, nil
}

func (pv PlacaVehiculo) String() string {
	return pv.value
}
