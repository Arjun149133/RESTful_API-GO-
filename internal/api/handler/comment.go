package handler

import (
	"example/restapi/internal/model"
	"example/restapi/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	Service service.CommentService
}

func (h *CommentHandler) CreateComment(c *gin.Context) {
	postId := c.Param("postId")
	var comment model.Comment

	comment.PostID = postId
	author, exists := c.Get("Author")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	comment.AuthorID = author.(map[string]interface{})["authorID"].(string)

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.Service.CreateComment(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comment)
}

func (h *CommentHandler) GetAllComments(c *gin.Context) {
	postId := c.Param("postId")

	comments, err := h.Service.GetAllComments(postId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (h *CommentHandler) GetCommentById(c *gin.Context) {
	commentId := c.Param("commentId")

	comment, err := h.Service.GetCommentById(commentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, comment)
}

func (h *CommentHandler) UpdateComment(c *gin.Context) {
	postId := c.Param("postId")

	var comment model.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.Service.UpdateComment(&comment, postId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comment)
}

func (h *CommentHandler) DeleteComment(c *gin.Context) {
	commentId := c.Param("commentId")
	id, err := strconv.ParseUint(commentId, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.DeleteComment(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "comment deleted"})
}
