package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Category struct {
	bun.BaseModel `bun:"table:categories"`

	ID        int64     `bun:",pk,autoincrement"`
	UserID    int64     `bun:",notnull" json:"user_id"`
	Name      string    `bun:"type:varchar(50),notnull" json:"category_name"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`

	User *User `bun:"rel:belongs-to,join:user_id=id"`
}
