package services

import (
	"database/sql"
	"errors"
	"final_project/internal/models"
	"final_project/internal/repository"
)

// получить всенх кандидатов
func GetAllCandidates(db *sql.DB) ([]models.Candidate, error) {
	return repository.GetAllCandidates(db)
}

// создать нового кандидата с базовой валидацией
func CreateCandidate(db *sql.DB, c *models.Candidate) error {
	if c.Name == "" || c.Position == "" || c.Email == "" || c.Phone == "" {
		return errors.New("all fields are required")
	}
	return repository.CreateCandidate(db, c)
}

// обновить кандидата с проверкой ID
func UpdateCandidate(db *sql.DB, c *models.Candidate) error {
	if c.ID == 0 {
		return errors.New("invalid candidate ID")
	}
	if c.Name == "" || c.Position == "" || c.Email == "" || c.Phone == "" {
		return errors.New("all fields are required")
	}
	return repository.UpdateCandidate(db, c)
}

// удалить кандидата
func DeleteCandidate(db *sql.DB, id int) error {
	if id == 0 {
		return errors.New("invalid candidate ID")
	}
	return repository.DeleteCandidate(db, id)
}
