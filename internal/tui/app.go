package app

import (
// "github.com/0b10headedcalf/daileet/internal/data"
// f "fmt"
// "os"
// tea "charm.land/bubbletea/v2"
)

type problemDifficulty int

const (
	diffEasy problemDifficulty = iota
	diffMed
	diffHard
)

type problem struct {
	name       string
	pattern    string
	url        string
	difficulty problemDifficulty
}

type model struct {
}

func main() {

	var difficulty = map[problemDifficulty]string{
		diffEasy: "Easy",
		diffMed:  "Medium",
		diffHard: "Hard",
	}
}
