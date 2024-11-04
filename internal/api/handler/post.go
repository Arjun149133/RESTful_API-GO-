package handler

import (
	"example/restapi/internal/model"
	"example/restapi/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	Service service.PostService
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var post model.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.Service.CreatePost(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, post)
}

func (h *PostHandler) GetAllPosts(c *gin.Context) {
	posts, err := h.Service.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}
	c.JSON(http.StatusOK, posts)
}
