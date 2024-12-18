package main

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin"       // Framework web
	"github.com/sirupsen/logrus"     // Librería para logging
)

func main() {
	// Configurar Logrus
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	log.Info("Iniciando la aplicación")

	// Crear una nueva instancia de Gin
	r := gin.Default()

	// Ruta simple
	r.GET("/", func(c *gin.Context) {
		log.Info("Se recibió una solicitud en /")
		c.JSON(200, gin.H{
			"message": "¡Hola, Mundo!",
		})
	})

	// Ruta para obtener el estado de la aplicación
	r.GET("/healthz", func(c *gin.Context) {
		log.Info("Comprobando el estado de la aplicación")
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Iniciar el servidor en el puerto 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("No se pudo iniciar el servidor: %v", err)
	}
}
