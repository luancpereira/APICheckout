package main

import (
	"github.com/luancpereira/APICheckout/apis/public/docs"
	"github.com/luancpereira/APICheckout/apis/public/server"
	"github.com/luancpereira/APICheckout/core/database"
	"github.com/luancpereira/APICheckout/core/errors"
)

func init() {
	errors.Factory{}.Start()
	database.Config{}.Start()

	docs.SwaggerInfo.Host = "localhost:9000"
}

//	@title			API Checkout
//	@version		1.0
//	@description	api checkout

// main entrypoint application
func main() {
	server.NewServer().Start()
}
