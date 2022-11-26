package domain

import (
	"time"
)

type QuestionNew struct {
	Body          string `json:"body"`
	OptionA       string `json:"option_a"`
	OptionB       string `json:"option_b"`
	OptionC       string `json:"option_c"`
	OptionD       string	`json:"option_d"`
	CorrectAnswer string `json:"correct_answer"`
	Score         int `json:"score"`
	CreatorID     int `json:"creator_id"`
	ExerciseID    int `json:"exercise_id"`
}

type Question struct {
	ID            int `json:"id"`
	Body          string `json:"body"`
	OptionA       string `json:"option_a"`
	OptionB       string `json:"option_b"`
	OptionC       string `json:"option_c"`
	OptionD       string	`json:"option_d"`
	CorrectAnswer string `json:"correct_answer"`
	Score         int `json:"score"`
	CreatorID     int `json:"creator_id"`
	ExerciseID    int `json:"exercise_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func NewQuestion(body, option_a, option_b, option_c, option_d, correct_answer string, exercise_id, score, creator_id int) (*Question, error) {
	return &Question{
		Body:          body,
		OptionA:       option_a,
		OptionB:       option_b,
		OptionC:       option_c,
		OptionD:       option_d,
		CorrectAnswer: correct_answer,
		Score:         score,
		CreatorID:     creator_id,
		ExerciseID:    exercise_id,
	}, nil
}


