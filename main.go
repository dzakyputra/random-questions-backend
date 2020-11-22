package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	controllers "github.com/dzakyputra/random-questions-backend/controllers"
	models "github.com/dzakyputra/random-questions-backend/models"
)

func main() {

	r := gin.Default()
	db := models.SetupModels()

	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/questions", controllers.FindQuestions)
	r.GET("/questions/:id", controllers.FindQuestionByID)

	r.Run()
}
