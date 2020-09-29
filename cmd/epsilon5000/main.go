package main

import (
	"awesomeProject/cmd/epsilon5000/apis"
	"awesomeProject/cmd/epsilon5000/config"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "gopkg.in/goracle.v2"
)

func main() {

	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	fmt.Println(config.Config.ZbxHost,config.Config.ZbxLogin,  config.Config.ZbxPassword)


	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("/")
	{
		v1.POST("Services", apis.GetSumService)
		v1.POST("Camera", apis.GetCameraOfZabbix)
	}

	r.Run(fmt.Sprintf(":%v", config.Config.ServerPort))


}
