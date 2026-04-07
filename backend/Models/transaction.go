package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Transaction struct {
	bun.BaseModel `bun:"table:transactions"`

	ID          int64     `bun:",pk,autoincrement" json:"id"`
	UserID      int64     `bun:",notnull" json:"user_id"`
	CategoryID  int64     `bun:",notnull" json:"category_id"`
	Amount      float64   `bun:"type:decimal(12,2),notnull" json:"amount"`
	Type        string    `bun:"type:enum('income','expense'),notnull" json:"type"`
	Description string    `bun:"type:text" json:"description"`
	Date        time.Time `bun:"type:date,notnull" json:"date"`
	CreatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`

	User     *User     `bun:"rel:belongs-to,join:user_id=id"`
	Category *Category `bun:"rel:belongs-to,join:category_id=id"`
}
