package scheduler

import (
	"time"

	"github.com/0b10headedcalf/daileet/internal/models"
)

// ProblemWithSchedule augments a Problem with computed scheduling info.
type ProblemWithSchedule struct {
	models.Problem
	DaysUntilDue int
}

// ComputeSchedule wraps a problem with scheduling metadata.
func ComputeSchedule(p models.Problem) ProblemWithSchedule {
	ps := ProblemWithSchedule{Problem: p}
	if p.DueDate == nil {
		ps.DaysUntilDue = 0
		return ps
	}
	ps.DaysUntilDue = int(p.DueDate.Sub(time.Now()).Hours() / 24)
	return ps
}
