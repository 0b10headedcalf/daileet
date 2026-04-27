package storage

import (
	"database/sql"
	"time"

	"github.com/0b10headedcalf/daileet/internal/models"
)

func scanProblems(rows *sql.Rows) ([]models.Problem, error) {
	var out []models.Problem
	for rows.Next() {
		var p models.Problem
		var due, reviewed sql.NullTime
		err := rows.Scan(&p.ID, &p.Title, &p.TitleSlug, &p.Difficulty, &p.Pattern, &p.URL,
			&p.Interval, &p.Repetitions, &p.EaseFactor, &due, &reviewed, &p.Status)
		if err != nil {
			return nil, err
		}
		if due.Valid {
			p.DueDate = &due.Time
		}
		if reviewed.Valid {
			p.LastReviewed = &reviewed.Time
		}
		out = append(out, p)
	}
	return out, rows.Err()
}

const selectCols = `SELECT id, title, title_slug, difficulty, pattern, url, interval, repetitions, ease_factor, due_date, last_reviewed, status FROM problems `

func ListDue(db *sql.DB) ([]models.Problem, error) {
	rows, err := db.Query(selectCols+`WHERE due_date IS NULL OR due_date <= ? ORDER BY CASE WHEN due_date IS NULL THEN 0 ELSE 1 END, due_date ASC, pattern, title`, time.Now())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanProblems(rows)
}

func ListNext(db *sql.DB) ([]models.Problem, error) {
	rows, err := db.Query(selectCols+`WHERE due_date IS NOT NULL AND due_date > ? ORDER BY due_date ASC, pattern, title`, time.Now())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanProblems(rows)
}

func ListSolved(db *sql.DB) ([]models.Problem, error) {
	rows, err := db.Query(selectCols + `WHERE repetitions > 0 ORDER BY last_reviewed DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanProblems(rows)
}

func ListAll(db *sql.DB) ([]models.Problem, error) {
	rows, err := db.Query(selectCols + `ORDER BY pattern, title`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanProblems(rows)
}

func GetProblemByID(db *sql.DB, id int) (models.Problem, error) {
	var p models.Problem
	var due, reviewed sql.NullTime
	err := db.QueryRow(selectCols+`WHERE id = ?`, id).Scan(
		&p.ID, &p.Title, &p.TitleSlug, &p.Difficulty, &p.Pattern, &p.URL,
		&p.Interval, &p.Repetitions, &p.EaseFactor, &due, &reviewed, &p.Status)
	if err != nil {
		return p, err
	}
	if due.Valid {
		p.DueDate = &due.Time
	}
	if reviewed.Valid {
		p.LastReviewed = &reviewed.Time
	}
	return p, nil
}

func UpdateProblemReview(db *sql.DB, p models.Problem) error {
	_, err := db.Exec(`
		UPDATE problems
		SET interval = ?, repetitions = ?, ease_factor = ?, due_date = ?, last_reviewed = ?, status = ?
		WHERE id = ?
	`, p.Interval, p.Repetitions, p.EaseFactor, p.DueDate, p.LastReviewed, p.Status, p.ID)
	return err
}

func DeleteProblemByID(db *sql.DB, id int) error {
	_, err := db.Exec(`DELETE FROM problems WHERE id = ?`, id)
	return err
}

func InsertProblem(db *sql.DB, p models.Problem) error {
	_, err := db.Exec(`
		INSERT INTO problems (title, title_slug, difficulty, pattern, url, interval, repetitions, ease_factor, due_date, last_reviewed, status)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, p.Title, p.TitleSlug, p.Difficulty, p.Pattern, p.URL, p.Interval, p.Repetitions, p.EaseFactor, p.DueDate, p.LastReviewed, p.Status)
	return err
}
