// Copyright (c) AppDynamics, Inc., and its affiliates
// 2023
// All Rights Reserved
// THIS IS UNPUBLISHED PROPRIETARY CODE OF APPDYNAMICS, INC.
// The copyright notice above does not evidence any actual or intended publication of such source code

package controller

import (
	"net/http"
	"player-service/model"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type IPlayerService interface{
	GetAll() ([]model.Player)
	GetById(id string) (player model.Player, ok bool)
}

type PlayerController struct {
	log                 *zap.SugaredLogger
	router              *gin.Engine
	playerService		IPlayerService
}

func NewPlayerController(logger *zap.SugaredLogger, playerService IPlayerService, router *gin.Engine) *PlayerController {
	controller := &PlayerController{
		log:                 logger,
		router:              router,
		playerService:       playerService,
	}

	// Ping test
	controller.router.GET("/ping", controller.healthApiHandler)

	v1 := controller.router.Group("v1/api/players")
	{
		v1.GET("/", controller.getAllPlayers)
		v1.GET("/:id", controller.getPlayerById)
	}

	return controller
}

func (controller *PlayerController) healthApiHandler(ctx *gin.Context) {
	controller.log.Info("Ping called")
	ctx.String(http.StatusOK, "pong")
}

func (controller *PlayerController) getAllPlayers(ctx *gin.Context) {
	controller.log.Info("Getting all players")
	players := controller.playerService.GetAll()

	ctx.JSON(http.StatusOK, players)

}

func (controller *PlayerController) getPlayerById(ctx *gin.Context) {
	id := ctx.Param("id")
	controller.log.Infow("Getting player by id", 
		"id", id,
	)
	player, ok := controller.playerService.GetById(id)
	if !ok {
		ctx.String(http.StatusNotFound, "Given i does not exist")
		return
	}
	ctx.JSON(http.StatusOK, player)
}

