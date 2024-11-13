package app

import (
	"errors"

	"github.com/snassiff/volcanoes/internal/domain"
)

// VolcanoService provee métodos para manejar la lógica de negocio de Volcano
type VolcanoService struct {
	repo domain.VolcanoRepository
}

// NewVolcanoService crea una nueva instancia de VolcanoService
func NewVolcanoService(repo domain.VolcanoRepository) *VolcanoService {
	return &VolcanoService{repo: repo}
}

func (s *VolcanoService) GetVolcanoes() ([]domain.Volcano, error) {
	return s.repo.FindAll()
}

func (s *VolcanoService) GetVolcanoByID(id uint) (*domain.Volcano, error) {
	return s.repo.FindByID(id)
}

func (s *VolcanoService) CreateVolcano(volcano *domain.Volcano) error {
	if volcano.Nombre == "" {
		return errors.New("nombre es requerido")
	}
	return s.repo.Create(volcano)
}

func (s *VolcanoService) UpdateVolcano(volcano *domain.Volcano) error {
	return s.repo.Update(volcano)
}

func (s *VolcanoService) DeleteVolcano(id uint) error {
	return s.repo.Delete(id)
}
