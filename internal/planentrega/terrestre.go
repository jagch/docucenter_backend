package planentrega

func NewPETerrestre(createPETerrestreReq PlanEntregaTerrestreReqRes) (PETerrestre, error) {

	idVO, err := NewIDTerrestre(createPETerrestreReq.ID)
	if err != nil {
		return PETerrestre{}, err
	}

	idClienteVO, err := NewIDCliente(createPETerrestreReq.IDCliente)
	if err != nil {
		return PETerrestre{}, err
	}

	tipoProductoVO, err := NewTipoProducto(createPETerrestreReq.TipoProducto)
	if err != nil {
		return PETerrestre{}, err
	}

	cantidadproductoVO, err := NewCantidadProducto(createPETerrestreReq.CantidadProducto)
	if err != nil {
		return PETerrestre{}, err
	}

	fechaRegistroVO, err := NewFechaRegistro(createPETerrestreReq.FechaRegistro)
	if err != nil {
		return PETerrestre{}, err
	}

	fechaEntregaVO, err := NewFechaEntrega(createPETerrestreReq.FechaEntrega)
	if err != nil {
		return PETerrestre{}, err
	}

	nroGuiaVO, err := NewNroGuia(createPETerrestreReq.NroGuia)
	if err != nil {
		return PETerrestre{}, err
	}

	idBodegaEntregaVO, err := NewIDBodegaEntrega(createPETerrestreReq.IDBodegaEntrega)
	if err != nil {
		return PETerrestre{}, err
	}

	precioEnvioVO, err := NewPrecioEnvio(createPETerrestreReq.PrecioEnvio)
	if err != nil {
		return PETerrestre{}, err
	}

	placaVehiculoVO, err := NewPlacaVehiculo(createPETerrestreReq.PlacaVehiculo)
	if err != nil {
		return PETerrestre{}, err
	}

	dsctoVO, err := NewDscto(createPETerrestreReq.Dscto)
	if err != nil {
		return PETerrestre{}, err
	}

	return PETerrestre{
		ID:              idVO,
		IDBodegaEntrega: idBodegaEntregaVO,
		PlacaVehiculo:   placaVehiculoVO,

		PlanEntrega: PlanEntrega{
			IDCliente:        idClienteVO,
			TipoProducto:     tipoProductoVO,
			CantidadProducto: cantidadproductoVO,
			FechaRegistro:    fechaRegistroVO,
			FechaEntrega:     fechaEntregaVO,
			PrecioEnvio:      precioEnvioVO,
			NroGuia:          nroGuiaVO,
			Dscto:            dsctoVO,
		},
	}, nil
}
