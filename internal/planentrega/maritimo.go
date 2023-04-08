package planentrega

func NewPEMaritimo(createPEMaritimoReq PlanEntregaMaritimoReqRes) (PEMaritimo, error) {

	idVO, err := NewIDMaritimo(createPEMaritimoReq.ID)
	if err != nil {
		return PEMaritimo{}, err
	}

	idClienteVO, err := NewIDCliente(createPEMaritimoReq.IDCliente)
	if err != nil {
		return PEMaritimo{}, err
	}

	tipoProductoVO, err := NewTipoProducto(createPEMaritimoReq.TipoProducto)
	if err != nil {
		return PEMaritimo{}, err
	}

	cantidadproductoVO, err := NewCantidadProducto(createPEMaritimoReq.CantidadProducto)
	if err != nil {
		return PEMaritimo{}, err
	}

	fechaRegistroVO, err := NewFechaRegistro(createPEMaritimoReq.FechaRegistro)
	if err != nil {
		return PEMaritimo{}, err
	}

	fechaEntregaVO, err := NewFechaEntrega(createPEMaritimoReq.FechaEntrega)
	if err != nil {
		return PEMaritimo{}, err
	}

	nroGuiaVO, err := NewNroGuia(createPEMaritimoReq.NroGuia)
	if err != nil {
		return PEMaritimo{}, err
	}

	idPuertoEntregaVO, err := NewIDPuertoEntrega(createPEMaritimoReq.IDPuertoEntrega)
	if err != nil {
		return PEMaritimo{}, err
	}

	precioEnvioVO, err := NewPrecioEnvio(createPEMaritimoReq.PrecioEnvio)
	if err != nil {
		return PEMaritimo{}, err
	}

	nroFlotaVO, err := NewNroFlota(createPEMaritimoReq.NroFlota)
	if err != nil {
		return PEMaritimo{}, err
	}

	dsctoVO, err := NewDscto(createPEMaritimoReq.Dscto)
	if err != nil {
		return PEMaritimo{}, err
	}

	return PEMaritimo{
		ID:              idVO,
		IDPuertoEntrega: idPuertoEntregaVO,
		NroFlota:        nroFlotaVO,

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
