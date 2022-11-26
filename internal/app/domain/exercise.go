package domain

import (
	"errors"
)

type Exercise struct {
	ID          int
	Title       string     `json:"judul"`
	Description string     `json:"description"`
	Questions   []Question `json:"questions"`
}

type ExerciseNew struct {
	ID          int
	Title       string
	Description string
}

// type Answer struct {
// 	ID         int
// 	ExerciseID int
// 	QuestionID int
// 	UserID     int
// 	Answer     string
// 	CreatedAt  time.Time
// 	UpdatedAt  time.Time
// }

func NewExercise(title, description string) (*Exercise, error) {
	if title == "" {
		return nil, errors.New("title is required")
	}

	if description == "" {
		return nil, errors.New("email is required")
	}

	return &Exercise{
		Title:       title,
		Description: description,
	}, nil
}
