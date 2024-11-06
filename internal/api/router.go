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

	authorRepo := &postgres.AuthorRepo{DB: db}
	authorService := &service.AuthorService{Repo: authorRepo}
	authorHandler := &handler.AuthorHandler{Service: *authorService}

	commentRepo := &postgres.CommentRepository{DB: db}
	commentService := &service.CommentService{Repo: commentRepo}
	commentHandler := &handler.CommentHandler{Service: *commentService}

	r.POST("/posts", postHandler.CreatePost)
	r.GET("/posts", postHandler.GetAllPosts)
	r.GET("/posts/:postId", postHandler.GetPostById)
	r.PUT("/posts/:postId", postHandler.UpdatePost)
	r.DELETE("/posts/:postId", postHandler.DeletePost)

	r.POST("/authors", authorHandler.CreateAuthor)
	r.GET("/authors", authorHandler.GetAllAuthors)
	r.GET("/authors/:authorId", authorHandler.GetAuthorById)
	r.PUT("/authors/:authorId", authorHandler.UpdateAuthor)
	r.DELETE("/authors/:authorId", authorHandler.DeleteAuthor)

	r.POST("/posts/:postId/comments", commentHandler.CreateComment)
	r.GET("/posts/:postId/comments", commentHandler.GetAllComments)
	r.GET("/posts/:postId/comments/:commentId", commentHandler.GetCommentById)
	r.PUT("/posts/:postId/comments/:commentId", commentHandler.UpdateComment)
	r.DELETE("/posts/:postId/comments/:commentId", commentHandler.DeleteComment)

	// r.POST("/temp", func(c *gin.Context) {
	// 	var temp model.Temp
	// 	if err := c.ShouldBindJSON(&temp); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": err,
	// 		})
	// 		return
	// 	}

	// 	if err := db.Create(temp).Error; err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, temp)
	// })

	return r
}
