// internal/app/service_test.go
package app

import (
	"testing"
	"time"

	"github.com/snassiff/volcanoes/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock del repositorio de volcanes
type MockVolcanoRepository struct {
	mock.Mock
}

func (m *MockVolcanoRepository) FindAll() ([]domain.Volcano, error) {
	args := m.Called()
	return args.Get(0).([]domain.Volcano), args.Error(1)
}

func (m *MockVolcanoRepository) FindByID(id uint) (*domain.Volcano, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Volcano), args.Error(1)
}

func (m *MockVolcanoRepository) Create(volcano *domain.Volcano) error {
	args := m.Called(volcano)
	return args.Error(0)
}

func (m *MockVolcanoRepository) Update(volcano *domain.Volcano) error {
	args := m.Called(volcano)
	return args.Error(0)
}

func (m *MockVolcanoRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

// Test para GetVolcanoes
func TestGetVolcanoes(t *testing.T) {
	mockRepo := new(MockVolcanoRepository)
	service := NewVolcanoService(mockRepo)

	volcanoes := []domain.Volcano{
		{ID: 1, Nombre: "Volcán Test"},
	}
	mockRepo.On("FindAll").Return(volcanoes, nil)

	result, err := service.GetVolcanoes()

	assert.NoError(t, err)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, "Volcán Test", result[0].Nombre)
	mockRepo.AssertExpectations(t)
}

// Test para CreateVolcano
func TestCreateVolcano(t *testing.T) {
	mockRepo := new(MockVolcanoRepository)
	service := NewVolcanoService(mockRepo)

	volcano := &domain.Volcano{
		Nombre:              "Nuevo Volcán",
		FechaUltimaErupcion: time.Now(),
	}

	mockRepo.On("Create", volcano).Return(nil)

	err := service.CreateVolcano(volcano)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
