package router

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go_prisma/controller"
	"net/http"
)

func NewRouter(postController *controller.PostController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Println(w, "Welcome!")
	})
	router.GET("/api/post/", postController.FindAll)
	router.GET("/api/post/:postId", postController.FindById)
	router.POST("/api/post/", postController.Create)
	router.PUT("/api/post/:postId", postController.Update)
	router.DELETE("/api/post/:postId", postController.Delete)

	return router
}
