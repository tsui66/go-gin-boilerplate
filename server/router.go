package server

import (
	"github.com/gin-gonic/gin"
	"github.com/quinton/go-gin-boilerplate/controller"
	"github.com/graphql-go/graphql"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	ping := new(controller.PingController)
	router.GET("/ping", ping.Status)
	
	return router
}