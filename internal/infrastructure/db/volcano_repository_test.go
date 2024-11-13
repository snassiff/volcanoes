// internal/infrastructure/db/volcano_repository_test.go
package db

import (
	"testing"
	"time"

	"github.com/snassiff/volcanoes/internal/domain"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Test para crear y encontrar un volcán
func TestGormVolcanoRepository(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	repo := NewGormVolcanoRepository(db)
	db.AutoMigrate(&domain.Volcano{})

	volcano := domain.Volcano{
		Nombre:              "Volcán Test",
		FechaUltimaErupcion: time.Now(),
	}

	// Test de creación
	err = repo.Create(&volcano)
	assert.NoError(t, err)

	// Test de búsqueda
	result, err := repo.FindByID(volcano.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Volcán Test", result.Nombre)
}
