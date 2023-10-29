package main

import (
	"os"
	"player-service/controller"
	"player-service/domain"
	"player-service/repository"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger := zap.NewExample().Sugar()
	defer logger.Sync()
	
	// initialize env variables
	port := os.Getenv("PORT")
	csvFilePath := os.Getenv("CSV_FILE_PATH")
	logger.Infow("Starting service with following env variables",
		"PORT", port,
		"CSV_FILE_PATH", csvFilePath,
	)

	// intialize all singleton components
	var (
		playerStore = repository.NewPlayerCsvStore(logger, csvFilePath) // TODO from conf
		playerService = domain.NewPlayerService(logger, playerStore)
		router = gin.Default()
		_ = controller.NewPlayerController(logger, playerService, router)
	)

	logger.Info("Player service is starting on port :8080")
	router.Run(":" + port)
}