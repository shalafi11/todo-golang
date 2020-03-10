package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"todo-golang/models"
)

func main() {
	fmt.Println("Hello World")
	r := gin.Default()
	fmt.Println("Gin default => ", r)

	testTodo := models.NewTodo("testTitle", "testBody")
	fmt.Println("Imported Todo: ", testTodo)

	testUser := models.NewUser("testName", "testEmail", "testPhone")
	fmt.Println("Imported User: ", testUser)

}
