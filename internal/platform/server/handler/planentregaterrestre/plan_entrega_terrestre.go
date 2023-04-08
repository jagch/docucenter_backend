package planentregaterrestre

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

func CreateHandler(peTerrestreStorage pentrega.PlanEntregaStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json; charset=utf-8")

		var req pentrega.PlanEntregaTerrestreReqRes

		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.CustomResponse{
				Data:    nil,
				Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
				Estado:  false,
			})

			return
		}
		req.ID = uuid.New().String()

		peTerrestre, err := pentrega.NewPETerrestre(req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.CustomResponse{
				Data:    nil,
				Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
				Estado:  true,
			})

			return
		}

		id, err := peTerrestreStorage.Create(ctx, peTerrestre)
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
			Mensaje: "plan de entrega terrestre creado con éxito",
			Estado:  true,
		})
	}
}

func GetAllHandler(peTerrestreStorage pentrega.PlanEntregaStorage, clienteStorage cliente.ClienteStorage) gin.HandlerFunc {
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

		peTerrestres, err := peTerrestreStorage.GetAll(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.CustomResponse{
				Data:    nil,
				Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
				Estado:  false,
			})
			return
		}

		peterr, err := setClientesbyIdInPeTerrestres(clientes, peTerrestres)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.CustomResponse{
				Data:    nil,
				Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
				Estado:  false,
			})
			return
		}

		ctx.JSON(http.StatusOK, models.CustomResponse{
			Data:    peterr,
			Mensaje: "planes de entrega terrestre obtenidos con éxito",
			Estado:  true,
		})
	}
}

func EditHandler(peTerrestreStorage pentrega.PlanEntregaStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json; charset=utf-8")

		var req pentrega.PlanEntregaTerrestreReqRes

		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.CustomResponse{
				Data:    nil,
				Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
				Estado:  false,
			})

			return
		}

		req.ID = ctx.Param("id")

		peTerrestre, err := pentrega.NewPETerrestre(req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.CustomResponse{
				Data:    nil,
				Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
				Estado:  true,
			})

			return
		}

		id, err := peTerrestreStorage.Update(ctx, peTerrestre)
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
			Mensaje: "plan de entrega terrestre actualizado con éxito",
			Estado:  true,
		})
	}
}

func DeleteHandler(peTerrestreStorage pentrega.PlanEntregaStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json; charset=utf-8")

		id := ctx.Param("id")

		err := peTerrestreStorage.Delete(ctx, id)
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
			Mensaje: "plan de entrega terrestre eliminado con éxito",
			Estado:  true,
		})
	}
}

func SearchHanlder(peTerrestreStorage pentrega.PlanEntregaStorage, clienteStorage cliente.ClienteStorage) gin.HandlerFunc {
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

		peTerrestres, err := peTerrestreStorage.Search(ctx, mapSearchParams, pageInt)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.CustomResponse{
				Data:    nil,
				Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
				Estado:  false,
			})
			return
		}

		peterr, err := setClientesbyIdInPeTerrestres(clientes, peTerrestres)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.CustomResponse{
				Data:    nil,
				Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
				Estado:  false,
			})
			return
		}

		msg := ""
		peTerr, _ := peterr.(pentrega.PlanesEntregaTerrestresReqRes)
		if len(peTerr) == 0 {
			msg = "No se encontraon resultados para la búsqueda"
		} else {
			msg = "búsqueda solicitada realizada con éxito"
		}

		ctx.JSON(http.StatusOK, models.CustomResponse{
			Data:    peterr,
			Mensaje: msg,
			Estado:  true,
		})
	}
}

func setClientesbyIdInPeTerrestres(clientes cliente.ClientesResponse, peTerrestres any) (any, error) {
	peTerr, ok := peTerrestres.(pentrega.PlanesEntregaTerrestresReqRes)
	if !ok {
		return nil, fmt.Errorf("%s", "error interno al cruzar información con clientes")
	}

	var planesEntregaTerrestreReqRes pentrega.PlanesEntregaTerrestresReqRes
	for _, pet := range peTerr {
		for _, c := range clientes {
			if pet.IDCliente == c.ID {
				pet.Cliente = c.Nombre
			}
		}
		planesEntregaTerrestreReqRes = append(planesEntregaTerrestreReqRes, pet)
	}

	return planesEntregaTerrestreReqRes, nil
}

func fillMapSearchParams(ctx *gin.Context) (map[string]pentrega.SearchParams, error) {
	mapSearchParams := make(map[string]pentrega.SearchParams)

	idCliente := ctx.Query("idCliente")
	fechaEntrega := ctx.Query("fechaEntrega")
	placaVehiculo := ctx.Query("placaVehiculo")

	if (idCliente == "") && (fechaEntrega == "") && (placaVehiculo == "") {
		return nil, fmt.Errorf("%s", "parámetros incorrectos, al menos ingrese o seleccione un parámetro para su búsqueda")
	}

	idClienteActive := false
	fechaEntregaActive := false
	placaVehiculoActive := false

	if idCliente != "" {
		idClienteActive = true
	}
	mapSearchParams[gb.FieldIDCliente] = pentrega.SearchParams{NameParam: gb.FieldIDCliente, Value: idCliente, Active: idClienteActive}

	if fechaEntrega != "" {
		fechaEntregaActive = true
	}
	mapSearchParams[gb.FieldFechaEntrega] = pentrega.SearchParams{NameParam: gb.FieldFechaEntrega, Value: fechaEntrega, Active: fechaEntregaActive}

	if placaVehiculo != "" {
		placaVehiculoActive = true
	}
	mapSearchParams[gb.FieldPlacaVehiculo] = pentrega.SearchParams{NameParam: gb.FieldPlacaVehiculo, Value: placaVehiculo, Active: placaVehiculoActive}

	return mapSearchParams, nil
}
