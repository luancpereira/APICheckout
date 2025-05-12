package config

import "os"

var (
	SERVER_PORT              = os.Getenv("SERVER_PORT")
	JWT_SECRET               = os.Getenv("JWT_SECRET")
	EXPIRATION_TOKEN_DEFAULT = os.Getenv("EXPIRATION_TOKEN_DEFAULT")
	SWAGGER_SERVER_HOST      = os.Getenv("SWAGGER_SERVER_HOST")
)
