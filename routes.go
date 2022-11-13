package main

import (
	"net/http"

	"github.com/aarthikrao/timeMachine/handlers/rest"
	"github.com/aarthikrao/timeMachine/process/client"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitTimeMachineHttpServer(
	cp *client.ClientProcess,
	log *zap.Logger,
) *http.Server {
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(gin.Recovery())
	// gin.SetMode(gin.ReleaseMode)

	// Health handler
	r.GET("/health", func(c *gin.Context) {
		// Return status ok
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Job handlers
	jrh := rest.CreateJobRestHandler(cp, log)
	job := r.Group("/job")
	{
		job.GET("/:collection/:jobID", jrh.GetJob)
		job.POST("/:collection", jrh.SetJob)
		job.DELETE("/:collection/:jobID", jrh.DeleteJob)
	}

	srv := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	return srv
}
