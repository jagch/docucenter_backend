package models

type CustomResponse struct {
	Data    any    `json:"data"`
	Mensaje string `json:"mensaje"`
	Estado  bool   `json:"estado"`
}
