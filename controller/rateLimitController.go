// Copyright (c) AppDynamics, Inc., and its affiliates
// 2023
// All Rights Reserved
// THIS IS UNPUBLISHED PROPRIETARY CODE OF APPDYNAMICS, INC.
// The copyright notice above does not evidence any actual or intended publication of such source code

package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RateLimitController struct {
	log                 *zap.SugaredLogger
	router              *gin.Engine
}

func NewRateLimitController(logger *zap.SugaredLogger, router *gin.Engine) *RateLimitController {
	controller := &RateLimitController{
		log:                 logger,
		router:              router,
	}

	controller.router.GET("/api/sleep", controller.rateLimiter(), controller.sleepApiHandler)

	return controller
}

func (controller *RateLimitController) sleepApiHandler(ctx *gin.Context) {
	controller.log.Info("Sleep called")
	ctx.String(http.StatusOK, "OK")
}

func (controller *RateLimitController) rateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		controller.log.Info("Received sleep call. Waiting for 5 secs.")

		time.Sleep(5 * time.Second)

		c.Next()
	}
}