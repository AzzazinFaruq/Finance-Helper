package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Transaction struct {
	bun.BaseModel `bun:"table:transactions"`

	ID          int64     `bun:",pk,autoincrement"`
	UserID      int64     `bun:",notnull"`
	CategoryID  int64     `bun:",notnull"`
	Amount      float64   `bun:"type:decimal(12,2),notnull"`
	Type        string    `bun:"type:enum('income','expense'),notnull"`
	Description string    `bun:"type:text"`
	Date        time.Time `bun:"type:date,notnull"`
	CreatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp"`

	User     *User     `bun:"rel:belongs-to,join:user_id=id"`
	Category *Category `bun:"rel:belongs-to,join:category_id=id"`
}
