package service

import (
	"context"

	"github.com/elwandyTD/api_bersama_umkm/model/web"
)

type AdminService interface {
	Create(ctx context.Context, request web.AdminCreateRequest) web.AdminResponse
}
