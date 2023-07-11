package main

import (
	"github.com/Johnfor93/go-payment/controllers"

	"github.com/gin-gonic/gin"
)

func main(){
	route := gin.Default()

	route.POST("/api/payment/create", paymentcontroller.Create)

	route.Run()
}