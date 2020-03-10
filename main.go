package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")
	r := gin.Default()
	fmt.Println("Gin default => ", r)
}
