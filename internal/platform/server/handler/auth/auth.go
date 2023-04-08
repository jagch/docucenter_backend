package auth

import (
	"fmt"

	"jagch/backend/internal/auth"
	"jagch/backend/internal/models"
	"jagch/backend/internal/platform/config"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthHandler(authUseCase auth.AuthUsecase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req auth.AuthRequest
		var authResponse auth.AuthResponse

		if err := ctx.BindJSON(&req); err != nil {

			ctx.JSON(http.StatusBadRequest, models.CustomResponse{
				Data:    nil,
				Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
				Estado:  false,
			})

			return
		}

		auth, err := auth.NewAuth(req.Usuario, req.Clave)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.CustomResponse{
				Data:    nil,
				Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
				Estado:  false,
			})

			return
		}

		id, err := authUseCase.Auth(ctx, auth)
		if id == 0 {
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, models.CustomResponse{
					Data:    nil,
					Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
					Estado:  false,
				})

				return
			}

			ctx.JSON(http.StatusOK, models.CustomResponse{
				Data:    nil,
				Mensaje: "el usuario y/o clave son incorrectas",
				Estado:  true,
			})

			return
		}

		// gen token
		token, err := createJWT(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.CustomResponse{
				Data:    nil,
				Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
				Estado:  false,
			})

			return
		}

		authResponse.Token = token

		ctx.JSON(http.StatusOK, models.CustomResponse{
			Data:    authResponse,
			Mensaje: "login successfully",
			Estado:  true,
		})
	}
}

func createJWT(id int) (string, error) {
	token := generoJWT()

	claims := token.Claims.(jwt.MapClaims)
	claims["id_user"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenSigned, err := token.SignedString([]byte(config.Config("TOP_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenSigned, nil
}

func generoJWT() *jwt.Token {
	return jwt.New(jwt.SigningMethodHS256)
}
