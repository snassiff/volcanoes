package main

import (
	"github.com/snassiff/volcanoes/internal/app"
	"github.com/snassiff/volcanoes/internal/app/handler"
	"github.com/snassiff/volcanoes/internal/infrastructure/config"
	"github.com/snassiff/volcanoes/internal/infrastructure/db"

	"github.com/gin-gonic/gin"
)

func main() {
	dbConn := config.GetDB()
	myEnv := config.NewEnv()
	myEnv.Env()

	// dbConn.AutoMigrate(&db.Volcano{}) // Auto-migración para la tabla Volcano, solamente al correr la primera vez

	repo := db.NewGormVolcanoRepository(dbConn)
	service := app.NewVolcanoService(repo)
	volcanoHandler := handler.NewVolcanoHandler(service)

	router := gin.Default()

	if myEnv.AppMode == "QUERY" {
		router.GET("/volcanoes", volcanoHandler.GetVolcanoes)
		router.GET("/volcanoes/:id", volcanoHandler.GetVolcano)
	} else if myEnv.AppMode == "COMMAND" {
		router.POST("/volcanoes", volcanoHandler.CreateVolcano)
	} else {
		panic("ERROR: No se seleccionó un modo de app válido, no se puede iniciar")
	}
	router.Run(":" + myEnv.ServerPort) // Inicia el servidor en el puerto configurado enb el archivo .env
}
