package domain

import (
	"time"
)

type Admin struct {
	Id       string    `json:"id"`
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
