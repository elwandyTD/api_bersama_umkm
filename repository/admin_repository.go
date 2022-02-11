package repository

import (
	"context"
	"database/sql"

	"github.com/elwandyTD/api_bersama_umkm/model/domain"
)

type AdminRepository interface {
	Create(ctx context.Context, tx *sql.Tx, admin domain.Admin) domain.Admin
}
