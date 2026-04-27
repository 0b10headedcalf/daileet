package scheduler

import "time"

// ReviewResult maps keybinds to SM-2 quality scores.
type ReviewResult int

const (
	Again ReviewResult = 1
	Hard  ReviewResult = 2
	Good  ReviewResult = 3
	Easy  ReviewResult = 4
)

// SM2State holds the scheduling parameters for a problem.
type SM2State struct {
	Interval    float64
	Repetitions int
	EaseFactor  float64
}

// DefaultSM2 returns the initial SM-2 state.
func DefaultSM2() SM2State {
	return SM2State{Interval: 0, Repetitions: 0, EaseFactor: 2.5}
}

// Grade maps user-friendly results to SM-2 quality q (0-5).
func (r ReviewResult) Quality() int {
	switch r {
	case Again:
		return 0
	case Hard:
		return 3
	case Good:
		return 4
	case Easy:
		return 5
	default:
		return 3
	}
}

// Apply runs the SM-2 algorithm and returns the next due date.
func (s *SM2State) Apply(result ReviewResult) time.Time {
	q := result.Quality()
	now := time.Now()

	if q < 3 {
		s.Repetitions = 0
		s.Interval = 1
	} else {
		switch s.Repetitions {
		case 0:
			s.Interval = 1
		case 1:
			s.Interval = 6
		default:
			s.Interval = s.Interval * s.EaseFactor
		}
		s.Repetitions++
	}

	s.EaseFactor = s.EaseFactor + (0.1 - (5-float64(q))*(0.08+(5-float64(q))*0.02))
	if s.EaseFactor < 1.3 {
		s.EaseFactor = 1.3
	}

	return now.Add(time.Duration(s.Interval*24) * time.Hour)
}
