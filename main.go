package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title`
	Status string `json:"status`
}

var todos = map[string]*Todo{}

func getTodosHandler(c *gin.Context) {

	todoVal := []*Todo{}

	for _, td := range todos {
		todoVal = append(todoVal, td)
	}

	c.JSON(http.StatusOK, todoVal)
}

func postTodosHandler(c *gin.Context) {

	todoVal := Todo{}

	if err := c.ShouldBindJSON(&todoVal); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	genId := len(todos)
	genId++
	id := strconv.Itoa(genId)
	todoVal.ID = id
	todos[id] = &todoVal

	c.JSON(http.StatusCreated, todoVal)
}

func getTodosByIDHandler(c *gin.Context) {

	id := c.Param("id")

	todoVal, ok := todos[id]

	if !ok {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	c.JSON(http.StatusOK, todoVal)
}

func putTodosHandler(c *gin.Context) {

	id := c.Param("id")

	todoVal := todos[id]

	if err := c.ShouldBindJSON(todoVal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todoVal)
}

func deleteTodosHandler(c *gin.Context) {

	id := c.Param("id")

	delete(todos, id)

	c.JSON(http.StatusOK, gin.H{"status": "deleted success"})

}

func main() {
	r := gin.Default()
	api := r.Group("/api")

	api.GET("/todos", getTodosHandler)
	api.POST("/todos", postTodosHandler)
	api.GET("/todos/:id", getTodosByIDHandler)
	api.PUT("/todos/:id", putTodosHandler)
	api.DELETE("/todos/:id", deleteTodosHandler)

	r.Run(":1234") //Go run listen and serve on 0.0.0.0:1234
}
