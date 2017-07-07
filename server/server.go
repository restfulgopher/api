package server

import (
	"log"

	"gopkg.in/gin-gonic/gin.v1"
)

// Serve triggers the server initialization
// serving on :8000 by default
func Serve(addr string) {
	if err := serverEngine().Run(addr); err != nil {
		log.Fatalf("failed to start server on port '%s': %s", addr, err)
	}
}

func serverEngine() *gin.Engine {
	r := gin.Default()

	// Register resource handlers
	v1 := r.Group("/" + apiVersion())
	v1.GET("/validate/:iban", validateHandler)
	return r
}

func apiVersion() string {
	return "v1"
}
