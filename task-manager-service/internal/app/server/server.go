package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func StartServer(eng *gin.Engine, port string, log *zap.Logger) {
	eng.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT"},
		AllowHeaders:    []string{"Content-Type", "Authorization"},
	}))
	log.Sugar().Infof("Starting http server on %s", port)
	if err := http.ListenAndServe(":"+port, eng); err != nil {
		log.Sugar().Fatalf("Error in server: %v", err)
	}

}
