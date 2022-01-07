package main

import (
	"net/http"

	ct "github.com/TT-EL/webshop/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/login", ct.Login)

	http.ListenAndServe(":8080", router)
}
