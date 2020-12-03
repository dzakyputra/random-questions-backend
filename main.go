package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	controllers "github.com/dzakyputra/random-questions-backend/controllers"
	models "github.com/dzakyputra/random-questions-backend/models"
)

func main() {

	r := gin.Default()

	// Provide CORS support
	r.Use(cors.Default())

	// Provide db variable to controllers
	db := models.SetupModels()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/questions", controllers.FindQuestions)
	r.GET("/question/:id", controllers.FindQuestionByID)
	r.GET("/questions/random", controllers.FindRandomQuestions)

	r.Run()
}
