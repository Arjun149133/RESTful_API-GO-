package api

import (
	"example/restapi/internal/api/handler"
	"example/restapi/internal/api/middleware"
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

	r.POST("/posts", middleware.AuthMiddleware(), postHandler.CreatePost)
	r.GET("/posts", postHandler.GetAllPosts)
	r.GET("/posts/:postId", postHandler.GetPostById)
	r.PUT("/posts/:postId", middleware.AuthMiddleware(), postHandler.UpdatePost)
	r.DELETE("/posts/:postId", middleware.AuthMiddleware(), postHandler.DeletePost)

	r.POST("/authors/register", authorHandler.CreateAuthor)
	r.POST("/authors/login", authorHandler.LoginAuthor)
	r.GET("/authors", authorHandler.GetAllAuthors)
	r.GET("/authors/:authorId", authorHandler.GetAuthorById)
	// r.GET("/authors/:authorId/posts", authorHandler.GetAuthorById)
	// r.GET("/authors/:authorId/comments", middleware.AuthMiddleware(), authorHandler.GetAuthorById)
	r.PUT("/authors/:authorId", middleware.AuthMiddleware(), authorHandler.UpdateAuthor)
	r.DELETE("/authors/:authorId", authorHandler.DeleteAuthor)

	r.POST("/posts/:postId/comments", middleware.AuthMiddleware(), commentHandler.CreateComment)
	r.GET("/posts/:postId/comments", commentHandler.GetAllComments)
	r.GET("/posts/:postId/comments/:commentId", commentHandler.GetCommentById)
	r.PUT("/posts/:postId/comments/:commentId", middleware.AuthMiddleware(), commentHandler.UpdateComment)
	r.DELETE("/posts/:postId/comments/:commentId", middleware.AuthMiddleware(), commentHandler.DeleteComment)

	return r
}
