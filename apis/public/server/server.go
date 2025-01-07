package server

import (
	"github.com/gin-gonic/gin"
	commonsConfig "github.com/luancpereira/APICheckout/apis/commons/config"
	commonsServer "github.com/luancpereira/APICheckout/apis/commons/server"
	"github.com/luancpereira/APICheckout/apis/public/server/routes"
)

type Server struct {
	Port   string
	Router *gin.Engine
}

func NewServer() (s Server) {

	s.Port = commonsConfig.SERVER_PORT
	s.Router = gin.Default()

	commonsServer.Server{}.SetupCORS(s.Router)
	commonsServer.Server{}.SetupSwagger(s.Router)
	s.setupRouterV1()

	return
}

func (s Server) Start() {
	address := ":" + s.Port
	err := s.Router.Run(address)
	if err != nil {
		panic(err)
	}
}

func (s Server) setupRouterV1() {
	freeRoutes := s.Router.Group("")

	checkout := routes.Checkout{}
	freeRoutes.GET("/api/checkout/transactions", checkout.GetList)
	freeRoutes.POST("/api/checkout", checkout.InsertTransaction)

}
