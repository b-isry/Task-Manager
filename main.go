package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type task struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	DueDate time.Time `json:"due_date"`
 	Status  string    `json:"status"`
}

var tasks = []task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
    {ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
    {ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},

}

func getTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}

func getTasksByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range tasks {
		if a.ID == id{
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message" : "task not found"})
}

func postTasks(c *gin.Context) {
	var newtask task

	if err := c.BindJSON(&newtask); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tasks = append(tasks, newtask)
	c.IndentedJSON(http.StatusCreated, gin.H{"message" : "Task created"})

}

func delTasks(c *gin.Context) {
	id := c.Param("id")

	for i, a := range tasks {
		if a.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message" : "Task removed" })
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message" : "task not found"})
}

func putTasks(c *gin.Context) {

	id := c.Param("id")

	var updatedTask task

	if err := c.ShouldBindJSON(&updatedTask); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	for i, a := range tasks {
		if a.ID == id {
			if updatedTask.Title != "" {
				tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				tasks[i].Description = updatedTask.Description
			}
			if updatedTask.Status != "" {
				tasks[i].Status = updatedTask.Status
			}
			c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
				return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message" : "task not found"})

}



func main() {

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/tasks", getTasks)
	router.GET("/tasks/:id", getTasksByID)
	router.PUT("/tasks/:id", putTasks)
	router.DELETE("/tasks/:id", delTasks)
	router.POST("/tasks", postTasks)
	router.Run()
}

