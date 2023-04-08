package cliente

type Cliente struct {
	ID     ClienteID
	Nombre ClienteNombre
}

type ClienteResponse struct {
	ID     string `json:"id"`
	Nombre string `json:"nombre"`
}

type ClientesResponse []ClienteResponse

type CreateClienteRequest struct {
	ID     string `json:"id"`
	Nombre string `json:"nombre"`
}
