package cliente

import (
	"fmt"
	"jagch/backend/internal/cliente"
	"jagch/backend/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateHandler(clienteStorage cliente.ClienteStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req cliente.CreateClienteRequest

		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.CustomResponse{
				Data:    nil,
				Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
				Estado:  false,
			})

			return
		}

		req.ID = uuid.New().String()

		cliente, err := cliente.NewCliente(req.ID, req.Nombre)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.CustomResponse{
				Data:    nil,
				Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
				Estado:  true,
			})

			return
		}

		clienteResponse, err := clienteStorage.Create(ctx, cliente)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.CustomResponse{
				Data:    nil,
				Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
				Estado:  false,
			})
			return
		}

		ctx.JSON(http.StatusCreated, models.CustomResponse{
			Data:    clienteResponse,
			Mensaje: "cliente creado con éxito",
			Estado:  true,
		})
	}
}

func GetAllHandler(clienteStorage cliente.ClienteStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		clientesResponse, err := clienteStorage.GetAll(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.CustomResponse{
				Data:    nil,
				Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
				Estado:  false,
			})
			return
		}

		ctx.JSON(http.StatusCreated, models.CustomResponse{
			Data:    clientesResponse,
			Mensaje: "cliente obtenidos con éxito",
			Estado:  true,
		})
	}
}
