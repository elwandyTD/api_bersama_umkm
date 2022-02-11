package domain

import (
	"database/sql"
	"time"
)

type Article struct {
	Id                string         `json:"id"`
	IdArticleCategory string         `json:"id_article_category"`
	IdAdmin           sql.NullString `json:"id_admin"`
	Title             string         `json:"title"`
	Description       string         `json:"description"`
	CoverImage        string         `json:"cover_image"`
	CoverImageTitle   string         `json:"cover_image_title"`
	Active            string         `json:"active"`
	CreateAt          time.Time      `json:"create_at"`
	UpdateAt          time.Time      `json:"update_at"`
}
