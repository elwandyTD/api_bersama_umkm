package domain

import "time"

type ArticleCategory struct {
	Id       string    `json:"id"`
	IdAdmin  string    `json:"id_admin"`
	Name     string    `json:"name"`
	Slug     string    `json:"slug"`
	Active   string    `json:"active"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}
