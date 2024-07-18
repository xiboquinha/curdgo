package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/xiboquinha/curdgo/src/controller/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("nao deu pra le a variavel de ambiente kdash")
	}

	router := gin.Default() //da pra usar o new tb, mas o default gera dois middleware pa nois

	routes.InitRoutes(&router.RouterGroup)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
