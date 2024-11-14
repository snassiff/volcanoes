package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/joho/godotenv"
)

// Para usar reflection se usa el tag modes, restringiendo los posibles valores a uno de los listados.
// Panic si llega otro o no se configura
type Env struct {
	DbUser     string
	DbPassword string
	DbServer   string
	DbDriver   string `modes:"valid=postgresql,mariadb,mysql"`
	DbPort     string
	DbName     string
	SslMode    string
	ServerPort string
	AppMode    string `modes:"valid=QUERY,COMMAND"`
}

func NewEnv() *Env {
	return &Env{}
}

// ValidateStruct valida los campos de un struct basados en el tag personalizado `modes`.
func (c *Env) Validate() error {
	val := reflect.ValueOf(*c)

	// Comprobar que el parámetro de entrada es un struct
	if val.Kind() != reflect.Struct {
		return errors.New("ValidateStruct solo acepta structs")
	}

	// Recorrer los campos del struct
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		fieldValue := val.Field(i)

		// Obtener el valor del tag "modes"
		tagValue := field.Tag.Get("modes")
		if tagValue == "" {
			continue // Si no hay tag, no hacemos ninguna validación
		}

		// Extraer los valores válidos del tag
		validValues := parseModesTag(tagValue)
		if len(validValues) == 0 {
			continue
		}

		// Verificar si el valor del campo es uno de los valores permitidos
		valueStr := fieldValue.String()
		if !contains(validValues, valueStr) {
			return fmt.Errorf("el valor '%s' en el campo '%s' no es válido; los valores permitidos son %v", valueStr, field.Name, validValues)
		}
	}
	return nil
}

// parseModesTag extrae los valores válidos de un tag personalizado `modes`
func parseModesTag(tag string) []string {
	parts := strings.Split(tag, "=")
	if len(parts) != 2 {
		return nil
	}
	return strings.Split(parts[1], ",")
}

// contains verifica si un valor está en una lista de valores
func contains(validValues []string, value string) bool {
	for _, v := range validValues {
		if v == value {
			return true
		}
	}
	return false
}

// NewEnv crea una nueva instancia de Env con valores cargados de las variables de entorno
func (c *Env) Env() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	c.DbUser = os.Getenv("DB_USER")
	c.DbPassword = os.Getenv("DB_PASSWORD")
	c.DbServer = os.Getenv("DB_SERVER")
	c.DbDriver = os.Getenv("DB_DRIVER")
	c.DbPort = os.Getenv("DB_PORT")
	c.DbName = os.Getenv("DB_NAME")
	c.SslMode = os.Getenv("SSLMODE")
	c.ServerPort = os.Getenv("SERVER_PORT")
	c.AppMode = os.Getenv("APP_MODE")

	//Si el environment no es correcto, nos vamos a panic y la app no arranca
	if err := c.Validate(); err != nil {
		panic("ERROR: NO se puede arrancar la app: " + err.Error())
	}

}
