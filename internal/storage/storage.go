package storage

import (
	"github.com/lucasvieirap/Note-Taking-App/internal/models"
)

type Storage interface {
	Get() *models.Note
	GetAll() []models.Note
	Update(id uint, note models.Note) 
	Delete(id uint) 
	Create(note models.Note) (uint, error)
}

