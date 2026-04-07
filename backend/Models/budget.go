package models

import (
	"github.com/uptrace/bun"
)

type Budget struct {
	bun.BaseModel `bun:"table:budgets"`

	ID         int64   `bun:",pk,autoincrement" json:"id"`
	UserID     int64   `bun:",notnull" json:"user_id"`
	CategoryID int64   `bun:",notnull" json:"category_id"`
	Amount     float64 `bun:"type:decimal(12,2),notnull" json:"amount"`
	Month      string  `bun:"type:char(7),notnull" json:"month"` // 2026-04

	User     *User     `bun:"rel:belongs-to,join:user_id=id"`
	Category *Category `bun:"rel:belongs-to,join:category_id=id"`
}
