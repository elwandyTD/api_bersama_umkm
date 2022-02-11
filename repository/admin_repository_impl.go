package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/elwandyTD/api_bersama_umkm/helper"
	"github.com/elwandyTD/api_bersama_umkm/model/domain"
)

type AdminRepositoryImpl struct{}

func NewAdminRepository() AdminRepository {
	return &AdminRepositoryImpl{}
}

func (repo *AdminRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, admin domain.Admin) domain.Admin {
	id := helper.GenerateNewUUID()
	hashPassword := helper.GenerateHashPassword(admin.Password)
	active := 1
	timeNow := time.Now().Format(time.RFC3339)
	SQL := "INSERT INTO admins(id, username, email, password, full_name, phone, address, active, create_at, update_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"

	res, err := tx.ExecContext(ctx, SQL, id, admin.Username, admin.Email, hashPassword, admin.FullName, admin.Phone, admin.Address, active, timeNow, timeNow)
	helper.PanicIfError(err)

	rows, err := res.RowsAffected()
	helper.PanicIfError(err)

	if rows < 1 {
		return admin
	}

	admin.Id = id
	admin.Active = string(rune(active))
	admin.CreateAt = time.Now()
	admin.UpdateAt = time.Now()

	return admin
}
