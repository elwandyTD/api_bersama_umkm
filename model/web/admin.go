package web

import (
	"time"
)

type AdminResponse struct {
	Email    string    `json:"email"`
	Username string    `json:"username"`
	FullName string    `json:"full_name"`
	Password string    `json:"password"`
	Address  *string   `json:"address"`
	Phone    *string   `json:"phone"`
	Active   string    `json:"active"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type AdminCreateRequest struct {
	Email    string  `validate:"required" json:"email"`
	Username string  `validate:"required" json:"username"`
	FullName string  `validate:"required" json:"full_name"`
	Password string  `validate:"required" json:"password"`
	Address  *string `json:"address"`
	Phone    *string `json:"phone"`
}
