package controller

import (
	"github.com/julienschmidt/httprouter"
	"go_prisma/data/request"
	"go_prisma/data/response"
	"go_prisma/helper"
	"go_prisma/service"
	"net/http"
)

type PostController struct {
	PostService service.PostService
}

func NewPostController(postService service.PostService) *PostController {
	return &PostController{postService}
}

// Create function
func (controller *PostController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {

	postCreateRequest := request.PostCreateRequest{}
	helper.ReadRequestBody(requests, &postCreateRequest)

	controller.PostService.Create(requests.Context(), postCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}
	helper.WriteRequest(writer, webResponse)
}

// Update function
func (controller *PostController) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {

	postUpdateRequest := request.PostUpdateRequest{}
	helper.ReadRequestBody(requests, &postUpdateRequest)

	postId := params.ByName("postId")
	postUpdateRequest.Id = postId
	controller.PostService.Update(requests.Context(), postUpdateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}
	helper.WriteRequest(writer, webResponse)
}

// FindAll function
func (controller *PostController) FindAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {

	result := controller.PostService.FindAll(requests.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}
	helper.WriteRequest(writer, webResponse)
}

// FindById function
func (controller *PostController) FindById(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	postId := params.ByName("postId")
	result := controller.PostService.FindByID(requests.Context(), postId)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}
	helper.WriteRequest(writer, webResponse)
}

func (controller *PostController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	postId := params.ByName("postId")
	controller.PostService.Delete(requests.Context(), postId)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}
	helper.WriteRequest(writer, webResponse)
}
