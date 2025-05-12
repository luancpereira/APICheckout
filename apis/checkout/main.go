package main

import (
	"github.com/luancpereira/APICheckout/apis/checkout/docs"
	"github.com/luancpereira/APICheckout/apis/checkout/server"
	commonsConfig "github.com/luancpereira/APICheckout/apis/commons/config"
	"github.com/luancpereira/APICheckout/core/database"
	"github.com/luancpereira/APICheckout/core/errors"
)

func init() {
	errors.Factory{}.Start()
	database.Config{}.Start()

	docs.SwaggerInfo.Host = commonsConfig.SWAGGER_SERVER_HOST
}

//	@title						API Checkout
//	@version					1.0
//	@description				API checkout
//	@securityDefinitions.apikey	JWT
//	@in							header
//	@name						Authorization

// main entrypoint application
func main() {
	server.NewServer().Start()
}
