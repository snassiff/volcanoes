package domain

import "time"

// Volcano representa un volcán en el sistema. Uso para para dto
type Volcano struct {
	ID                  uint      `json:"ID"`
	Nombre              string    `json:"Nombre"`
	Descripcion         string    `json:"Descripcion"`
	Departamento        string    `json:"Departamento"`
	Latitud             float64   `json:"Latitud"`
	Longitud            float64   `json:"Longitud"`
	Altura              int       `json:"Altura"`
	Tipo                string    `json:"Tipo"`
	Activo              bool      `json:"Activo"`
	FechaUltimaErupcion time.Time `json:"FechaUltimaErupcion"`
	// CreatedAt           time.Time
	// UpdatedAt           time.Time
}
