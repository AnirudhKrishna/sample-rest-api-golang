package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	Id        string `json:"id`
	Item      string `json:"title`
	Completed bool   `json:"completed`
}

var todos = []todo{

	{Id: "1", Item: "Clean Room", Completed: true},
	{Id: "2", Item: "Read Book", Completed: false},
	{Id: "3", Item: "Record Video", Completed: true},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.Run("localhost:9090")
}
