package domain

import (
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// TestVolcanoModel prueba la creación y lectura del modelo Volcano
func TestVolcanoModel(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error al conectar a la base de datos en memoria: %v", err)
	}

	// Migrar el modelo Volcano
	if err := db.AutoMigrate(&Volcano{}); err != nil {
		t.Fatalf("Error al migrar el modelo Volcano: %v", err)
	}

	// Crear un volcán de prueba
	testVolcano := Volcano{
		Nombre:              "Volcán Test",
		Descripcion:         "Descripción de prueba",
		Departamento:        "Departamento de prueba",
		Latitud:             10.0,
		Longitud:            -10.0,
		Altura:              1500,
		Tipo:                "Estratovolcán",
		Activo:              true,
		FechaUltimaErupcion: time.Now(),
	}

	// Guardar el volcán en la base de datos
	if err := db.Create(&testVolcano).Error; err != nil {
		t.Fatalf("Error al crear el volcán: %v", err)
	}

	// Leer el volcán de la base de datos y verificar
	var volcano Volcano
	if err := db.First(&volcano, testVolcano.ID).Error; err != nil {
		t.Fatalf("Error al leer el volcán: %v", err)
	}
	if volcano.Nombre != "Volcán Test" {
		t.Errorf("Se esperaba 'Volcán Test', pero se obtuvo '%s'", volcano.Nombre)
	}
}
