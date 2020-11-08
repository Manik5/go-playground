package main

import (
  "net/http"
  "github.com/gin-gonic/gin"

  "github.com/Manik5/go_crash_course/mini-projects/rest-api/models"
  "github.com/Manik5/go_crash_course/mini-projects/rest-api/controllers"
)

func main() {
  r := gin.Default()

  models.ConnectDatabase()

  r.GET("/", func(c * gin.Context) {
    c.JSON(http.StatusOK, gin.H{"data": "hello world"})
  })

  r.GET("/books", controllers.FindBooks)

  r.Run()
}
