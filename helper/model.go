package helper

import (
	"github.com/elwandyTD/api_bersama_umkm/model/domain"
	"github.com/elwandyTD/api_bersama_umkm/model/web"
)

func ToAdminResponse(admin domain.Admin) web.AdminResponse {
	return web.AdminResponse{
		Email:    admin.Email,
		Username: admin.Username,
		FullName: admin.FullName,
		Password: admin.Password,
		Address:  admin.Address,
		Phone:    admin.Phone,
	}
}
