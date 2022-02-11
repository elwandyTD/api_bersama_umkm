package service

import (
	"context"
	"database/sql"

	"github.com/elwandyTD/api_bersama_umkm/helper"
	"github.com/elwandyTD/api_bersama_umkm/model/domain"
	"github.com/elwandyTD/api_bersama_umkm/model/web"
	"github.com/elwandyTD/api_bersama_umkm/repository"
	"github.com/go-playground/validator/v10"
)

type AdminServiceImpl struct {
	AdminRepository repository.AdminRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewAdminService(adminRepository repository.AdminRepository, DB *sql.DB, Validate *validator.Validate) AdminService {
	return &AdminServiceImpl{
		AdminRepository: adminRepository,
		DB:              DB,
		Validate:        Validate,
	}
}

func (service *AdminServiceImpl) Create(ctx context.Context, request web.AdminCreateRequest) web.AdminResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	admin := domain.Admin{
		Email:    request.Email,
		Username: request.Username,
		FullName: request.FullName,
		Password: request.Password,
		Address:  request.Address,
		Phone:    request.Phone,
	}

	admin = service.AdminRepository.Create(ctx, tx, admin)

	return helper.ToAdminResponse(admin)
}
