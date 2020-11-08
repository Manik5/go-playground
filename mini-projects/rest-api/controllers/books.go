package controllers

import (
  "github.com/gin-gonic/gin"
  "github.com/Manik5/go_crash_course/mini-projects/rest-api/models"
)

// GET /books
// Get all books

// FindBooks func
func FindBooks(c *gin.Context) {
  var books []models.Book
  models.DB.Find(&books)

  c.JSON(http.StatusOK, gin.H{"data": books})
}
