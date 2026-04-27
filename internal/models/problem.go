package models

import "time"

type Difficulty string

const (
	Easy   Difficulty = "Easy"
	Medium Difficulty = "Medium"
	Hard   Difficulty = "Hard"
)

type Problem struct {
	ID           int
	Title        string
	TitleSlug    string
	Difficulty   Difficulty
	Pattern      string
	URL          string
	Interval     float64
	Repetitions  int
	EaseFactor   float64
	DueDate      *time.Time
	LastReviewed *time.Time
	Status       string // new | learning | review
}
