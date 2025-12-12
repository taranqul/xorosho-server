package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func StartServer(eng *gin.Engine, port string, log *zap.Logger) {
	log.Sugar().Infof("Starting http server on %s", port)
	if err := http.ListenAndServe(":"+port, eng); err != nil {
		log.Sugar().Fatalf("Error in server: %v", err)
	}

}
