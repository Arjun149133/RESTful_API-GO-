package api

import (
	"example/restapi/internal/api/handler"
	"example/restapi/internal/repository/postgres"
	"example/restapi/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	postRepo := &postgres.PostRepo{DB: db}
	postService := &service.PostService{Repo: postRepo}
	postHandler := &handler.PostHandler{Service: *postService}

	r.POST("/posts", postHandler.CreatePost)
	r.GET("/posts", postHandler.GetAllPosts)
	r.GET("/posts/:postId", postHandler.GetPostById)
	r.PUT("/posts/:postId", postHandler.UpdatePost)
	r.DELETE("/posts/:postId", postHandler.DeletePost)

	return r
}
