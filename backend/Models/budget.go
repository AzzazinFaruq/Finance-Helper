package models

import (
	"github.com/uptrace/bun"
)

type Budget struct {
	bun.BaseModel `bun:"table:budgets"`

	ID         int64   `bun:",pk,autoincrement"`
	UserID     int64   `bun:",notnull"`
	CategoryID int64   `bun:",notnull"`
	Amount     float64 `bun:"type:decimal(12,2),notnull"`
	Month      string  `bun:"type:char(7),notnull"` // 2026-04

	User     *User     `bun:"rel:belongs-to,join:user_id=id"`
	Category *Category `bun:"rel:belongs-to,join:category_id=id"`
}
