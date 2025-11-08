package repository

import (
	"database/sql"
	"final_project/internal/models"
)

func GetAllCandidates(db *sql.DB) ([]models.Candidate, error) {
	query := `SELECT id, name, position, email, phone, created_at FROM candidates`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var candidates []models.Candidate
	for rows.Next() {
		var c models.Candidate
		if err := rows.Scan(&c.ID, &c.Name, &c.Position, &c.Email, &c.Phone, &c.CreatedAt); err != nil {
			return nil, err
		}
		candidates = append(candidates, c)
	}
	return candidates, nil
}

func CreateCandidate(db *sql.DB, c *models.Candidate) error {
	query := `INSERT INTO candidates (name, position, email, phone) VALUES ($1, $2, $3, $4) RETURNING id, created_at`
	return db.QueryRow(query, c.Name, c.Position, c.Email, c.Phone).Scan(&c.ID, &c.CreatedAt)
}

func UpdateCandidate(db *sql.DB, c *models.Candidate) error {
	query := `UPDATE candidates SET name=$1, position=$2, email=$3, phone=$4 WHERE id=$5 RETURNING created_at`
	return db.QueryRow(query, c.Name, c.Position, c.Email, c.Phone, c.ID).Scan((&c.CreatedAt))
}

func DeleteCandidate(db *sql.DB, id int) error {
	query := `DELETE FROM candidates WHERE id=$1`
	_, err := db.Exec(query, id)
	return err
}
