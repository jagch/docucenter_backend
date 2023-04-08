package auth

type Auth struct {
	Usuario AuthUsuario
	Clave   AuthClave
}

type AuthRequest struct {
	Usuario string `json:"usuario"`
	Clave   string `json:"clave"`
}

type AuthResponse struct {
	Token string `json:"token"`
}
