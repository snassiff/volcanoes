// internal/app/handler/volcano_test.go
package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/snassiff/volcanoes/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock del servicio VolcanoService
type MockVolcanoService struct {
	mock.Mock
}

func (m *MockVolcanoService) GetVolcanoes() ([]domain.Volcano, error) {
	args := m.Called()
	return args.Get(0).([]domain.Volcano), args.Error(1)
}

func (m *MockVolcanoService) GetVolcanoByID(id uint) (*domain.Volcano, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Volcano), args.Error(1)
}

func (m *MockVolcanoService) CreateVolcano(volcano *domain.Volcano) error {
	args := m.Called(volcano)
	return args.Error(0)
}

// TestGetVolcanoes verifica el controlador de obtener todos los volcanes
func TestGetVolcanoes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockVolcanoService)
	handler := NewVolcanoHandler(mockService)

	// Crear datos de prueba
	volcanoes := []domain.Volcano{
		{ID: 1, Nombre: "Volcán Test"},
	}
	mockService.On("GetVolcanoes").Return(volcanoes, nil)

	// Crear una solicitud de prueba
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Llamar al controlador
	handler.GetVolcanoes(c)

	// Verificar el resultado
	assert.Equal(t, http.StatusOK, w.Code)
	var response []domain.Volcano
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(response))
	assert.Equal(t, "Volcán Test", response[0].Nombre)
}

// TestGetVolcano verifica el controlador para obtener un volcán por ID
func TestGetVolcano(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockVolcanoService)
	handler := NewVolcanoHandler(mockService)

	// Crear datos de prueba
	volcano := &domain.Volcano{ID: 1, Nombre: "Volcán Test"}
	mockService.On("GetVolcanoByID", uint(1)).Return(volcano, nil)

	// Crear una solicitud de prueba
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	// Llamar al controlador
	handler.GetVolcano(c)

	// Verificar el resultado
	assert.Equal(t, http.StatusOK, w.Code)
	var response domain.Volcano
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Volcán Test", response.Nombre)
}

// TestCreateVolcano verifica el controlador para crear un volcán
func TestCreateVolcano(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockVolcanoService)
	handler := NewVolcanoHandler(mockService)

	// Crear un volcán de prueba
	volcano := domain.Volcano{
		Nombre:              "Nuevo Volcán",
		Descripcion:         "Descripción de prueba",
		Departamento:        "Antioquia",
		Latitud:             1.23,
		Longitud:            -77.02,
		Altura:              2500,
		Tipo:                "Estratovolcán",
		Activo:              true,
		FechaUltimaErupcion: "2022-01-01T00:00:00Z",
	}

	// Convertir el volcán a JSON
	body, err := json.Marshal(volcano)
	assert.NoError(t, err)

	// Configurar el mock para crear un volcán
	mockService.On("CreateVolcano", &volcano).Return(nil)

	// Crear una solicitud de prueba
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPost, "/volcanoes", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Request.Header.Set("User-Agent", "test-agent")

	// Llamar al controlador
	handler.CreateVolcano(c)

	// Verificar el resultado
	assert.Equal(t, http.StatusCreated, w.Code)
	var response domain.Volcano
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Nuevo Volcán", response.Nombre)
}

// TestCreateVolcanoInvalidData verifica el controlador para manejar datos inválidos al crear un volcán
func TestCreateVolcanoInvalidData(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockVolcanoService)
	handler := NewVolcanoHandler(mockService)

	// Crear un volcán de prueba con datos inválidos
	volcano := domain.Volcano{
		Nombre: "",
	}
	body, err := json.Marshal(volcano)
	assert.NoError(t, err)

	// Crear una solicitud de prueba con datos inválidos
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPost, "/volcanoes", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-Type", "application/json")

	// Llamar al controlador
	handler.CreateVolcano(c)

	// Verificar el resultado
	assert.Equal(t, http.StatusBadRequest, w.Code)
	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "errors")
}
