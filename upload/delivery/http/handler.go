package http

import (
	"net/http"
	"static/upload"

	gin "github.com/gin-gonic/gin"
)

type Handler struct {
	useCase upload.UseCase
}

func NewHandler(useCase upload.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) Load(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      "bad",
			"description": "Well managed to extract file from form",
		})
	}

	name, err := h.useCase.Load(file)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":      "bad",
			"description": "failed to upload file",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"name":   name,
	})
}

func (h *Handler) GetImage(c *gin.Context) {
	id := c.Param("id")

	c.File("files/" + id)
}

func (h *Handler) UploadPage(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", gin.H{})
}
