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

// FindRandomQuestion return a random question in the database
func FindRandomQuestion(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	var question models.Question
	var total int

	// Find total questions
	db.Model(&question).Count(&total)

	var id = rand.Intn(total-1) + 1

	// Find the question based on the random generated number
	if err := db.Find(&question, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": question})
}
