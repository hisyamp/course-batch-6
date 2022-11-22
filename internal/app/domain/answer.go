package domain

import (
	"time"
)

type Answer struct {
	Answer     string
	UserID     int
	QuestionID          int
	ExerciseID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}


func NewAnswer(answer string, user_id, question_id, exercise_id int, created_at, updated_at time.Time) (*Answer, error) {
	return &Answer{
		Answer:          answer,
		UserID:       user_id,
		QuestionID:       question_id,
		ExerciseID:       exercise_id,
		CreatedAt: created_at,
		UpdatedAt:       updated_at,
		
	}, nil
}


