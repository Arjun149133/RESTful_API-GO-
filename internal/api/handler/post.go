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

func (h *PostHandler) GetPostById(c *gin.Context) {
	postId := c.Param("postId")

	post, err := h.Service.GetPostById(postId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *PostHandler) UpdatePost(c *gin.Context) {
	postId := c.Param("postId")

	var updatedPost model.Post
	updatedPost.ID = postId
	if err := c.ShouldBindJSON(&updatedPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.Service.UpdatePost(&updatedPost); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedPost)
}

func (h *PostHandler) DeletePost(c *gin.Context) {
	postId := c.Param("postId")

	if err := h.Service.DeletePost(postId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "Deleted Post succesfully",
	})
}
