package main

import (
	"log"
	"template/service/db"
	"template/service/server"

	"github.com/gin-gonic/gin"
)

func main() {
	loadDatabase()
	router := gin.Default()
	server.Routes(router)
	log.Fatal(router.Run(":4747"))

}

func loadDatabase() {
	Database := db.Connect()
	Database.MustExec(db.Schema)
}
