package handler

import (
	"example/restapi/internal/config"
	"example/restapi/internal/model"
	"example/restapi/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthorHandler struct {
	Service service.AuthorService
}

func (h *AuthorHandler) CreateAuthor(c *gin.Context) {
	var author model.Author
	err := c.ShouldBindJSON(&author)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if author.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "password cannot be empty",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(author.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	author.Password = string(hashedPassword)

	if err := h.Service.CreateAuthor(&author); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, author)
}

func (h *AuthorHandler) LoginAuthor(c *gin.Context) {
	var author model.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	authorInDB, err := h.Service.LoginAuthor(author.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(authorInDB.Password), []byte(author.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid password",
		})
		return
	}
	cfg := config.Load()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": authorInDB.Email,
		"name":  authorInDB.Name,
	})

	// Sign and get the complete encoded token as a string using the secret
	accessToken, err := token.SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})
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

	author, err := h.Service.GetAuthorById(authorId)
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

	var updatedAuthor model.Author
	updatedAuthor.ID = authorId

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

	if err := h.Service.DeleteAuthor(authorId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "Deleted Author succesfully",
	})
}
