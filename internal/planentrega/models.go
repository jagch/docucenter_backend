package planentrega

import "time"

type PlanEntrega struct {
	IDCliente        IDCliente
	TipoProducto     TipoProducto
	CantidadProducto CantidadProducto
	FechaRegistro    FechaRegistro
	FechaEntrega     FechaEntrega
	PrecioEnvio      PrecioEnvio
	NroGuia          NroGuia
	Dscto            Dscto
}

type PETerrestre struct {
	ID              IDTerrestre
	IDBodegaEntrega IDBodegaEntrega
	PlacaVehiculo   PlacaVehiculo

	PlanEntrega PlanEntrega
}

type PETerrestres []PETerrestre

type PlanEntregaTerrestreReqRes struct {
	Key             int    `json:"key"`
	ID              string `json:"id"`
	IDBodegaEntrega int    `json:"id_bodega_entrega"`
	BodegaEntrega   string `json:"bodega_entrega"`
	PlacaVehiculo   string `json:"placa_vehiculo"`

	IDCliente        string    `json:"id_cliente"`
	Cliente          string    `json:"cliente"`
	TipoProducto     string    `json:"tipo_producto"`
	CantidadProducto int       `json:"cantidad_producto"`
	FechaRegistro    time.Time `json:"fecha_registro"`
	FechaEntrega     time.Time `json:"fecha_entrega"`
	PrecioEnvio      float64   `json:"precio_envio"`
	NroGuia          string    `json:"nro_guia"`
	Dscto            float64   `json:"dscto"`
}

type PlanesEntregaTerrestresReqRes []PlanEntregaTerrestreReqRes

/************************************************/
type PEMaritimo struct {
	ID              IDMaritimo
	IDPuertoEntrega IDPuertoEntrega
	NroFlota        NroFlota

	PlanEntrega PlanEntrega
}

type PEMaritimos []PEMaritimo

type PlanEntregaMaritimoReqRes struct {
	Key             int    `json:"key"`
	ID              string `json:"id"`
	IDPuertoEntrega int    `json:"id_puerto_entrega"`
	PuertoEntrega   string `json:"puerto_entrega"`
	NroFlota        string `json:"nro_flota"`

	IDCliente        string    `json:"id_cliente"`
	Cliente          string    `json:"cliente"`
	TipoProducto     string    `json:"tipo_producto"`
	CantidadProducto int       `json:"cantidad_producto"`
	FechaRegistro    time.Time `json:"fecha_registro"`
	FechaEntrega     time.Time `json:"fecha_entrega"`
	PrecioEnvio      float64   `json:"precio_envio"`
	NroGuia          string    `json:"nro_guia"`
	Dscto            float64   `json:"dscto"`
}

type PlanesEntregaMaritimosReqRes []PlanEntregaMaritimoReqRes

type SearchParams struct {
	NameParam string
	Value     any
	Active    bool
}

type SearchRequest struct {
	Rows     any `json:"rows"`
	Total    int `json:"total"`
	Page     int `json:"page"`
	LastPage int `json:"last_page"`
}
