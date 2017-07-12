package server

import (
	"log"

	"gopkg.in/gin-gonic/gin.v1"
)

// Serve triggers the server initialization
func Serve(addr string) {
	if err := serverEngine().Run(addr); err != nil {
		log.Fatalf("failed to start server on port '%s': %s", addr, err)
	}
}

func serverEngine() *gin.Engine {
	r := gin.Default()

	// Register resource handlers
	v1 := r.Group("/" + apiVersion())
	iban := v1.Group("")
	iban.GET("/iban/valid/:iban", validHandler)
	return r
}

func apiVersion() string {
	return "v1"
}
