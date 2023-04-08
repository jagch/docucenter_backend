package server

import (
	"fmt"
	"jagch/backend/internal/auth"
	"jagch/backend/internal/cliente"
	planentrega "jagch/backend/internal/planentrega"
	handlerAuth "jagch/backend/internal/platform/server/handler/auth"
	handlerCliente "jagch/backend/internal/platform/server/handler/cliente"
	handlerPlanEntregaMaritimo "jagch/backend/internal/platform/server/handler/planentregamaritimo"
	handlerPlanEntregaTerrestre "jagch/backend/internal/platform/server/handler/planentregaterrestre"
	"jagch/backend/internal/platform/server/middlewares"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	//deps
	clienteStorage     cliente.ClienteStorage
	authUseCase        auth.AuthUsecase
	peTerrestreStorage planentrega.PlanEntregaStorage
	peMaritimoStorage  planentrega.PlanEntregaStorage
}

func New(host, port string, authUseCase auth.AuthUsecase, clienteStorage cliente.ClienteStorage, peTerrestreStorage planentrega.PlanEntregaStorage, peMaritimoStorage planentrega.PlanEntregaStorage) Server {
	srv := Server{
		httpAddr: fmt.Sprintf("%s:%s", host, port),
		engine:   gin.New(),

		authUseCase:        authUseCase,
		clienteStorage:     clienteStorage,
		peTerrestreStorage: peTerrestreStorage,
		peMaritimoStorage:  peMaritimoStorage,
	}

	srv.registerRoutes()

	return srv
}

func (s *Server) Run() error {
	log.Println("server running on ", s.httpAddr)

	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	api := s.engine.Group("/api", middlewares.CORSMiddleware())
	apiV1 := api.Group("/v1")
	// auth: sin token
	apiV1.POST("/auth", handlerAuth.AuthHandler(s.authUseCase))

	// grupo docucenter: con token
	apiV1Docucenter := apiV1.Group("/docucenter")
	apiV1Docucenter.Use(middlewares.JwtAuthMiddleware())
	apiV1Docucenter.POST("/clientes", handlerCliente.CreateHandler(s.clienteStorage))
	apiV1Docucenter.GET("/clientes", handlerCliente.GetAllHandler(s.clienteStorage))
	apiV1Docucenter.POST("/peterrestres", handlerPlanEntregaTerrestre.CreateHandler(s.peTerrestreStorage))
	apiV1Docucenter.GET("/peterrestres", handlerPlanEntregaTerrestre.GetAllHandler(s.peTerrestreStorage, s.clienteStorage))
	apiV1Docucenter.PATCH("/peterrestres/:id", handlerPlanEntregaTerrestre.EditHandler(s.peTerrestreStorage))
	apiV1Docucenter.DELETE("/peterrestres/:id", handlerPlanEntregaTerrestre.DeleteHandler(s.peTerrestreStorage))
	apiV1Docucenter.GET("/peterrestres/search", handlerPlanEntregaTerrestre.SearchHanlder(s.peTerrestreStorage, s.clienteStorage))
	apiV1Docucenter.POST("/pemaritimos", handlerPlanEntregaMaritimo.CreateHandler(s.peMaritimoStorage))
	apiV1Docucenter.GET("/pemaritimos", handlerPlanEntregaMaritimo.GetAllHandler(s.peMaritimoStorage, s.clienteStorage))
	apiV1Docucenter.PATCH("/pemaritimos/:id", handlerPlanEntregaMaritimo.EditHandler(s.peMaritimoStorage))
	apiV1Docucenter.DELETE("/pemaritimos/:id", handlerPlanEntregaMaritimo.DeleteHandler(s.peMaritimoStorage))
	apiV1Docucenter.GET("/pemaritimos/search", handlerPlanEntregaMaritimo.SearchHanlder(s.peMaritimoStorage, s.clienteStorage))
}
