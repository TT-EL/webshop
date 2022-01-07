package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	name := c.DefaultQuery("name", "Tom")
	fmt.Println(name)

}
