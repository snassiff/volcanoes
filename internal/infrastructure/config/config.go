package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GetDB crea una conexión a la base de datos PostgreSQL usando GORM
// Las credenciaoles vienen del env para determinar si la instalcia es
// la de lectura o la de escritura
func GetDB() *gorm.DB {

	e := NewEnv()
	e.Env() //cargar todas las variables y validar

	dsn := "host=" + e.DbServer + " user=" + e.DbUser + " password=" + e.DbPassword + " dbname=" + e.DbName + " port=" + e.DbPort + " sslmode=" + e.SslMode + " TimeZone=America/Bogota"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	return db
}
