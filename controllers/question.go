package controllers

import (
	"math/rand"
	"net/http"

	models "github.com/dzakyputra/random-questions-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// FindQuestions is for return all of the questions
func FindQuestions(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	var questions []models.Question
	db.Find(&questions)
	c.JSON(http.StatusOK, gin.H{"data": questions})
}

// FindQuestionByID return a question by its ID
func FindQuestionByID(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var question models.Question
	if err := db.Find(&question, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": question})
}

// FindRandomQuestions return a random question in the database
func FindRandomQuestions(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	var questions []models.Question
	var total int

	// Find total questions
	db.Model(&questions).Count(&total)

	var totalRandom = 20
	var start = rand.Intn((total - totalRandom)) - 1
	var end = start + totalRandom

	// Find the question based on the random generated number
	db.Where("id >= ? AND id <= ?", start, end).Find(&questions)
	c.JSON(http.StatusOK, gin.H{"data": questions})
}
