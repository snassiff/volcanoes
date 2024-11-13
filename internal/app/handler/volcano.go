package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/snassiff/volcanoes/internal/app"
	"github.com/snassiff/volcanoes/internal/domain"
	u "github.com/snassiff/volcanoes/internal/utils"

	"github.com/gin-gonic/gin"
)

type VolcanoHandler struct {
	service *app.VolcanoService
}

// NewVolcanoHandler crea una nueva instancia de VolcanoHandler
func NewVolcanoHandler(service *app.VolcanoService) *VolcanoHandler {
	return &VolcanoHandler{service: service}
}

// GetVolcanoes maneja la solicitud para obtener todos los volcanes
func (h *VolcanoHandler) GetVolcanoes(c *gin.Context) {
	volcanoes, err := h.service.GetVolcanoes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, volcanoes)
}

// GetVolcano maneja la solicitud para obtener un volcán por ID
func (h *VolcanoHandler) GetVolcano(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	volcano, err := h.service.GetVolcanoByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Volcán no encontrado"})
		return
	}
	c.JSON(http.StatusOK, volcano)
}

// CreateVolcano maneja la solicitud para crear un volcán
func (h *VolcanoHandler) CreateVolcano(c *gin.Context) {
	var volcano domain.Volcano
	if err := c.ShouldBindJSON(&volcano); err != nil {
		// Verificar si el error es del tipo ValidationErrors de "validator"
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			// Recorrer la lista de errores de validación para obtener detalles
			errorMessages := make(map[string]string)
			for _, e := range validationErrors {
				// `e.Field()` proporciona el nombre del campo que causó el error
				errorMessages[e.Field()] = fmt.Sprintf("Campo '%s' es inválido: %s", e.Field(), e.ActualTag())
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
			return
		}

		// En caso de otro tipo de error de binding
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos - " + err.Error()})
		return
	}
	clientdata := u.ClientInfo{}
	if err := c.ShouldBindHeader(&clientdata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	} else {
		clientdata.ClientIP = c.ClientIP()
		log.Default().Println(clientdata) //Mostrar datos del navegador y cliente en el log de sistema (por hacer algo)
	}
	if err := h.service.CreateVolcano(&volcano); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, volcano)
}
