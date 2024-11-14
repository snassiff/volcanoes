package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GetDB crea una conexión a la base de datos PostgreSQL usando GORM
// Las credenciaoles vienen del env para determinar si la instalcia es
// la de lectura o la de escritura
func GetDB() *gorm.DB {
	var dsn string
	var db *gorm.DB
	var err error

	e := NewEnv()
	e.Env() //cargar todas las variables y validar

	if e.DbDriver == "postgresql" {
		dsn = "host=" + e.DbServer + " user=" + e.DbUser + " password=" + e.DbPassword + " dbname=" + e.DbName + " port=" + e.DbPort + " sslmode=" + e.SslMode + " TimeZone=America/Bogota"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else if e.DbDriver == "mysql" || e.DbDriver == "mariadb" {
		dsn = e.DbUser + ":" + e.DbPassword + "@tcp(" + e.DbServer + ":" + e.DbPort + ")/" + e.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else {
		log.Fatalf("No se soporta el driver %v", e.DbDriver)
	}

	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	return db
}
