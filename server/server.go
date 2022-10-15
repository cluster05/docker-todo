package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	err := InitConfig()
	if err != nil {
		log.Println(err)
	}

	client, err := InitMongoDB()
	if err != nil {
		log.Println(err)
	}

	router := gin.Default()
	router.Use(cors.Default())
	InitHandler(router, client)

	PORT := fmt.Sprintf(":%v", config.PORT)
	err = router.Run(PORT)
	if err != nil {
		log.Println(err)
	}

}
