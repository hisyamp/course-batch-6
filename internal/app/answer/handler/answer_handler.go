package handler

import (
	"exercise/internal/app/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AnswerHandler struct {
	db *gorm.DB
}

func NewAnswerHandler(db *gorm.DB) *AnswerHandler {
	return &AnswerHandler{
		db: db,
	}
}

func (ah AnswerHandler) CreateAnswer(c *gin.Context) {
	var createAnswer domain.Answer
	if err := c.ShouldBind(&createAnswer); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid body",
		})
	}


	answer, err := domain.NewAnswer(createAnswer.Answer,createAnswer.UserID,createAnswer.QuestionID,createAnswer.ExerciseID,createAnswer.CreatedAt,createAnswer.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	if err := ah.db.Create(answer).Error; err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, map[string]string{
		"Message": "Sukses Create Answer!",
	})

}
