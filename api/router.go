package api

import (
	"github.com/Riku0617/myapi/api/middlewares"
	"github.com/Riku0617/myapi/controllers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter(con *controllers.ArticleController, con2 *controllers.CommentController) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middlewares.LoggingMiddleware)

	r.Route("/article", func(r chi.Router) {
		r.Post("/", con.PostArticleHandler)
		r.Get("/list", con.ArticleListHandler)
		r.Get("/{articleID}", con.ArticleDetailHandler)
		r.Get("/nice", con.PostNiceHandler)
	})

	r.Post("/comment", con2.PostCommentHandler)

	return r
}
