package handler

import (
	"exercise/internal/app/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QuestionHandler struct {
	db *gorm.DB
}

func NewQuestionHandler(db *gorm.DB) *QuestionHandler {
	return &QuestionHandler{
		db: db,
	}
}

func (uh QuestionHandler) CreateQuestion(c *gin.Context) {
	var createQuestion domain.Question
	if err := c.ShouldBind(&createQuestion); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid body",
		})
	}

	// question, err := domain.NewQuestion(createQuestion.Body, createQuestion.OptionA, createQuestion.OptionB, createQuestion.OptionC, createQuestion.OptionD, createQuestion.CorrectAnswer, createQuestion.CreatorID, createQuestion.Score, createQuestion.ExerciseID)
	question, err := domain.NewQuestion(createQuestion.Body, createQuestion.OptionA, createQuestion.OptionB, createQuestion.OptionC, createQuestion.OptionD, createQuestion.CorrectAnswer, createQuestion.CreatorID, createQuestion.Score, createQuestion.ExerciseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	if err := uh.db.Create(question).Error; err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, map[string]string{
		"Message": "Sukses Create Question!",
	})

}
