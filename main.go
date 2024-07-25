package main

import (
	"github.com/gin-gonic/gin"
	"github.com/loco/conf"
	"github.com/loco/services"
)

func main() {
	r := gin.Default()

	services.InitServices(r)

	r.Run(conf.String("loco.port", ":9092"))
}
