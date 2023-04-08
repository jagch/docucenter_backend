package global

import "time"

var CurrentTime = time.Now().Format("2006.01.02 15:04:05")

var (
	FieldIDCliente     string = "id_cliente"
	FieldFechaEntrega  string = "fecha_entrega"
	FieldPlacaVehiculo string = "placa_vehiculo"
	FieldNroFlota      string = "nro_flota"
)

var (
	PerPage int = 5
)
