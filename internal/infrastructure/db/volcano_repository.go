package db

import (
	"time"

	"github.com/snassiff/volcanoes/internal/domain"

	"gorm.io/gorm"
)

// Estructura para gorm
type Volcano struct {
	ID                  uint      `gorm:"primaryKey;autoIncrement"`
	Nombre              string    `gorm:"size:100;not null"`
	Descripcion         string    `gorm:"type:text"`
	Departamento        string    `gorm:"size:100;not null"`
	Latitud             float64   `gorm:"not null"`
	Longitud            float64   `gorm:"not null"`
	Altura              int       `gorm:"not null"`
	Tipo                string    `gorm:"size:50;not null"`
	Activo              bool      `gorm:"not null"`
	FechaUltimaErupcion time.Time `gorm:"not null"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type GormVolcanoRepository struct {
	db *gorm.DB
}

// NewGormVolcanoRepository crea una nueva instancia de GormVolcanoRepository
func NewGormVolcanoRepository(db *gorm.DB) domain.VolcanoRepository {
	return &GormVolcanoRepository{db: db}
}

func (r *GormVolcanoRepository) FindAll() ([]domain.Volcano, error) {
	var volcanoes []domain.Volcano
	err := r.db.Find(&volcanoes).Error
	return volcanoes, err
}

func (r *GormVolcanoRepository) FindByID(id uint) (*domain.Volcano, error) {
	var volcano domain.Volcano
	err := r.db.First(&volcano, id).Error
	if err != nil {
		return nil, err
	}
	return &volcano, nil
}

func (r *GormVolcanoRepository) Create(volcano *domain.Volcano) error {
	return r.db.Create(volcano).Error
}

func (r *GormVolcanoRepository) Update(volcano *domain.Volcano) error {
	return r.db.Save(volcano).Error
}

func (r *GormVolcanoRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Volcano{}, id).Error
}
