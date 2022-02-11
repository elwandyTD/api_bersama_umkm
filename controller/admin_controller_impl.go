package controller

import (
	"net/http"

	"github.com/elwandyTD/api_bersama_umkm/helper"
	"github.com/elwandyTD/api_bersama_umkm/model/web"
	"github.com/elwandyTD/api_bersama_umkm/service"
	"github.com/julienschmidt/httprouter"
)

type AdminControllerImpl struct {
	AdminService service.AdminService
}

func NewAdminController(adminService service.AdminService) AdminController {
	return &AdminControllerImpl{
		AdminService: adminService,
	}
}

func (controller *AdminControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	adminCreateRequest := web.AdminCreateRequest{}
	helper.RequestDecode(request, &adminCreateRequest)

	adminResponse := controller.AdminService.Create(request.Context(), adminCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   adminResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
