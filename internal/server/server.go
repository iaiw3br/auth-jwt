package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"main/config"
	authError "main/internal/error"
	"main/internal/route"
)

func Run() error {
	if err := config.Init(); err != nil {
		log.Fatalf("Failed to read config: %+v", err)
	}

	router := gin.Default()
	route.New(router)

	address := viper.GetString("RUN_ADDRESS")
	port := viper.GetString("PORT")

	err := router.Run(address + port)
	if err != nil {
		log.Fatalf(authError.ErrFailedStartServer)
	}

	return nil
}
