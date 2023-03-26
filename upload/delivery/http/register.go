package http

import (
	"static/upload"

	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc upload.UseCase) {
	h := NewHandler(uc)

	heavyApiLoadEndpoints := router.Group("/")
	{
		heavyApiLoadEndpoints.GET("/", h.UploadPage)
		heavyApiLoadEndpoints.POST("/load", h.Load)
		heavyApiLoadEndpoints.GET("/:id", h.GetImage)
	}
}
