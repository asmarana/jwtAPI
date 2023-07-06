package main

import (
	"log"
	"template/service/server"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	server.Routes(router)
	log.Fatal(router.Run(":4747"))

}
