package server

import (
	"personnel-management-backend/config"
	"personnel-management-backend/router"

	"github.com/gin-gonic/gin"
)

func SetupServer() {
	cfg := config.Conf.Server
	r := gin.Default()
	router.SetupRouter(r)
	r.Run(cfg.Host + ":" + cfg.Port)
}
