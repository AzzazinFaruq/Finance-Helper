package models

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID        int64     `bun:",pk,autoincrement" json:"id"`
	Username  string    `bun:"type:varchar(50),notnull" json:"username"`
	Password  string    `bun:"type:varchar(100),notnull" json:"-"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"updated_at"`

	Categories   []*Category   `bun:"rel:has-many,join:id=user_id"`
	Transactions []*Transaction `bun:"rel:has-many,join:id=user_id"`

}