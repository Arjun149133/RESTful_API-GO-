package handler

import (
	"example/restapi/internal/model"
	"example/restapi/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthorHandler struct {
	Service service.AuthorService
}

func (h *AuthorHandler) CreateAuthor(c *gin.Context) {
	var author model.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.Service.CreateAuthor(&author); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, author)
}

func (h *AuthorHandler) GetAllAuthors(c *gin.Context) {
	authors, err := h.Service.Repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, authors)
}

func (h *AuthorHandler) GetAuthorById(c *gin.Context) {
	authorId := c.Param("authorId")

	id, err := strconv.ParseUint(authorId, 10, 32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	author, err := h.Service.GetAuthorById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, author)
}

func (h *AuthorHandler) UpdateAuthor(c *gin.Context) {
	authorId := c.Param("authorId")

	id, err := strconv.ParseUint(authorId, 10, 32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	var updatedAuthor model.Author
	updatedAuthor.ID = uint(id)
	if err := c.ShouldBindJSON(&updatedAuthor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.Service.UpdateAuthor(&updatedAuthor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedAuthor)
}

func (h *AuthorHandler) DeleteAuthor(c *gin.Context) {
	authorId := c.Param("authorId")
	id, err := strconv.ParseUint(authorId, 10, 32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := h.Service.DeleteAuthor(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "Deleted Author succesfully",
	})
}
