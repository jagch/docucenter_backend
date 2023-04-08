package planentregamaritimo

import (
	"fmt"
	"jagch/backend/internal/cliente"
	gb "jagch/backend/internal/global"
	"jagch/backend/internal/models"
	pentrega "jagch/backend/internal/planentrega"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateHandler(peMaritimoStorage pentrega.PlanEntregaStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json; charset=utf-8")

		var req pentrega.PlanEntregaMaritimoReqRes

		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.CustomResponse{
				Data:    nil,
				Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
				Estado:  false,
			})

			return
		}
		req.ID = uuid.New().String()

		peMaritimo, err := pentrega.NewPEMaritimo(req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.CustomResponse{
				Data:    nil,
				Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
				Estado:  true,
			})

			return
		}

		id, err := peMaritimoStorage.Create(ctx, peMaritimo)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.CustomResponse{
				Data:    nil,
				Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
				Estado:  false,
			})
			return
		}

		ctx.JSON(http.StatusCreated, models.CustomResponse{
			Data:    id,
			Mensaje: "plan de entrega maritimo creado con éxito",
			Estado:  true,
		})
	}
}

func GetAllHandler(peMaritimoStorage pentrega.PlanEntregaStorage, clienteStorage cliente.ClienteStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		clientes, err := clienteStorage.GetAll(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.CustomResponse{
				Data:    nil,
				Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
				Estado:  false,
			})
			return
		}

		peMaritimos, err := peMaritimoStorage.GetAll(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.CustomResponse{
				Data:    nil,
				Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
				Estado:  false,
			})
			return
		}

		pemari, err := setClientesbyIdInPeMaritimos(clientes, peMaritimos)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.CustomResponse{
				Data:    nil,
				Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
				Estado:  false,
			})
			return
		}

		ctx.JSON(http.StatusOK, models.CustomResponse{
			Data:    pemari,
			Mensaje: "planes de entrega terrestre obtenidos con éxito",
			Estado:  true,
		})
	}
}

func EditHandler(peMaritimoStorage pentrega.PlanEntregaStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json; charset=utf-8")

		var req pentrega.PlanEntregaMaritimoReqRes

		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.CustomResponse{
				Data:    nil,
				Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
				Estado:  false,
			})

			return
		}

		req.ID = ctx.Param("id")

		peMaritimo, err := pentrega.NewPEMaritimo(req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.CustomResponse{
				Data:    nil,
				Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
				Estado:  true,
			})

			return
		}

		id, err := peMaritimoStorage.Update(ctx, peMaritimo)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.CustomResponse{
				Data:    nil,
				Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
				Estado:  false,
			})
			return
		}

		ctx.JSON(http.StatusOK, models.CustomResponse{
			Data:    id,
			Mensaje: "plan de entrega maritimo actualizado con éxito",
			Estado:  true,
		})
	}
}

func DeleteHandler(peMaritimoStorage pentrega.PlanEntregaStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json; charset=utf-8")

		id := ctx.Param("id")

		err := peMaritimoStorage.Delete(ctx, id)
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
			Mensaje: "plan de entrega maritimo eliminado con éxito",
			Estado:  true,
		})
	}
}

func SearchHanlder(peMaritimoStorage pentrega.PlanEntregaStorage, clienteStorage cliente.ClienteStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json; charset=utf-8")

		mapSearchParams, err := fillMapSearchParams(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.CustomResponse{
				Data:    nil,
				Mensaje: fmt.Sprintf("solicitud errónea: %s", err.Error()),
				Estado:  false,
			})
			return
		}

		pageStr := ctx.Query("page")
		pageInt, err := strconv.Atoi(pageStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.CustomResponse{
				Data:    nil,
				Mensaje: fmt.Sprintf("el parámetro <page> no puede ser vacío o es incorrecto. %s", err.Error()),
				Estado:  false,
			})
			return
		}

		clientes, err := clienteStorage.GetAll(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.CustomResponse{
				Data:    nil,
				Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
				Estado:  false,
			})
			return
		}

		peMaritimos, err := peMaritimoStorage.Search(ctx, mapSearchParams, pageInt)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.CustomResponse{
				Data:    nil,
				Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
				Estado:  false,
			})
			return
		}

		pemari, err := setClientesbyIdInPeMaritimos(clientes, peMaritimos)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.CustomResponse{
				Data:    nil,
				Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
				Estado:  false,
			})
			return
		}

		msg := ""
		peMari, _ := pemari.(pentrega.PlanesEntregaMaritimosReqRes)
		if len(peMari) == 0 {
			msg = "no se encontraon resultados para la busqueda"
		} else {
			msg = "búsqueda solicitada realizada con éxito"
		}

		ctx.JSON(http.StatusOK, models.CustomResponse{
			Data:    peMari,
			Mensaje: msg,
			Estado:  true,
		})
	}
}

func setClientesbyIdInPeMaritimos(clientes cliente.ClientesResponse, peMaritimos any) (any, error) {
	peMari, ok := peMaritimos.(pentrega.PlanesEntregaMaritimosReqRes)
	if !ok {
		return nil, fmt.Errorf("%s", "error interno al cruzar información con clientes")
	}

	var planesEntregaMaritimosReqRes pentrega.PlanesEntregaMaritimosReqRes
	for _, pem := range peMari {
		for _, c := range clientes {
			if pem.IDCliente == c.ID {
				pem.Cliente = c.Nombre
			}
		}
		planesEntregaMaritimosReqRes = append(planesEntregaMaritimosReqRes, pem)
	}

	return planesEntregaMaritimosReqRes, nil
}

func fillMapSearchParams(ctx *gin.Context) (map[string]pentrega.SearchParams, error) {
	mapSearchParams := make(map[string]pentrega.SearchParams)

	idCliente := ctx.Query("idCliente")
	fechaEntrega := ctx.Query("fechaEntrega")
	nroFlota := ctx.Query("nroFlota")

	if (idCliente == "") && (fechaEntrega == "") && (nroFlota == "") {
		return nil, fmt.Errorf("%s", "parámetros incorrectos, al menos ingrese o seleccione un parámetro para su búsqueda")
	}

	idClienteActive := false
	fechaEntregaActive := false
	nroFlotaActive := false

	if idCliente != "" {
		idClienteActive = true
	}
	mapSearchParams[gb.FieldIDCliente] = pentrega.SearchParams{NameParam: gb.FieldIDCliente, Value: idCliente, Active: idClienteActive}

	if fechaEntrega != "" {
		fechaEntregaActive = true
	}
	mapSearchParams[gb.FieldFechaEntrega] = pentrega.SearchParams{NameParam: gb.FieldFechaEntrega, Value: fechaEntrega, Active: fechaEntregaActive}

	if nroFlota != "" {
		nroFlotaActive = true
	}
	mapSearchParams[gb.FieldNroFlota] = pentrega.SearchParams{NameParam: gb.FieldNroFlota, Value: nroFlota, Active: nroFlotaActive}

	return mapSearchParams, nil
}
