package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/luancpereira/APICheckout/apis/checkout/server/routes"
	commonsAuth "github.com/luancpereira/APICheckout/apis/commons/auth"
	commonsConfig "github.com/luancpereira/APICheckout/apis/commons/config"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	Port   string
	Router *gin.Engine
}

// Setups

func SetupCORS(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Accept", "Authorization"},
	}))
}

func SetupSwagger(router *gin.Engine) {
	router.GET("/docs/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// Setups

func NewServer() (s Server) {

	s.Port = commonsConfig.SERVER_PORT
	s.Router = gin.Default()

	SetupCORS(s.Router)
	SetupSwagger(s.Router)
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
	authRoutes := s.Router.Group("").Use(commonsAuth.JWT{}.Middleware())

	user := routes.User{}
	freeRoutes.POST("/api/users", user.PostUser)

	auth := routes.Auth{}
	freeRoutes.POST("/api/auth/login", auth.PostLogin)

	checkout := routes.Checkout{}
	authRoutes.POST("/api/checkout", checkout.InsertTransaction)
	authRoutes.GET("/api/checkout/transactions/country/:country", checkout.GetList)
	authRoutes.GET("/api/checkout/transactions/:transactionID/country/:country", checkout.GetByID)
	authRoutes.DELETE("/api/checkout/transactions/:transactionID", checkout.DeleteTransaction)
	authRoutes.PUT("/api/checkout/transactions", checkout.UpdateTransaction)

}
