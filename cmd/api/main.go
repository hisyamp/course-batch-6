package main

import (
	answerHandler "exercise/internal/app/answer/handler"
	"exercise/internal/app/database"
	"exercise/internal/app/exercise/handler"
	questionHandler "exercise/internal/app/question/handler"
	userHandler "exercise/internal/app/user/handler"
	"exercise/internal/pkg/middleware"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{

		v1.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, map[string]string{
				"message": "hello world",
			})
		})
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, map[string]string{
				"message": "Hi!",
			})
		})
	
		db := database.NewConnDatabase()
		exerciseHandler := handler.NewExerciseHandler(db)
		userHandler := userHandler.NewUserHandler(db)
		questionHandler := questionHandler.NewQuestionHandler(db)
		answerHandler := answerHandler.NewAnswerHandler(db)
		v1.GET("/exercises/:id", middleware.WithAuh(), exerciseHandler.GetExerciseByID)
		v1.GET("/exercises/:id/score", middleware.WithAuh(), exerciseHandler.GetScore)
	
		v1.POST("/register", userHandler.Register)
		v1.POST("/answer", answerHandler.CreateAnswer)
		v1.POST("/exercise", middleware.WithAuh(), exerciseHandler.CreateExercise)
		v1.POST("/question", middleware.WithAuh(), questionHandler.CreateQuestion)
		v1.POST("/login", userHandler.Login)
	}
	var port string
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	} else {
		port = "1234"
	}
	log.Fatal(r.Run(fmt.Sprintf(":%s", port)))
}
