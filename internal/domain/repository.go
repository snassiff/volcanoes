package domain

// VolcanoRepository define las operaciones de almacenamiento para Volcano
type VolcanoRepository interface {
	FindAll() ([]Volcano, error)
	FindByID(id uint) (*Volcano, error)
	Create(volcano *Volcano) error
	Update(volcano *Volcano) error
	Delete(id uint) error
}
