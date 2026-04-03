package Setup

import (
	"simple_crud/Models"
	"context"
	"log"
)

func Migrate() {
	ctx := context.Background()

	// create tables dulu
	models := []interface{}{
		(*models.User)(nil),
		(*models.Category)(nil),
		(*models.Transaction)(nil),
		(*models.Budget)(nil),
	}

	for _, model := range models {
		_, err := DB.NewCreateTable().
			Model(model).
			IfNotExists().
			Exec(ctx)

		if err != nil {
			log.Fatalf("Error creating table %T: %v", model, err)
		}
	}

	// add foreign keys
	_, _ = DB.Exec(`
		ALTER TABLE categories
		ADD CONSTRAINT fk_categories_user
		FOREIGN KEY (user_id) REFERENCES users(id)
		ON DELETE CASCADE;
	`)

	_, _ = DB.Exec(`
		ALTER TABLE transactions
		ADD CONSTRAINT fk_transactions_user
		FOREIGN KEY (user_id) REFERENCES users(id)
		ON DELETE CASCADE;
	`)

	_, _ = DB.Exec(`
		ALTER TABLE transactions
		ADD CONSTRAINT fk_transactions_category
		FOREIGN KEY (category_id) REFERENCES categories(id)
		ON DELETE CASCADE;
	`)

	_, _ = DB.Exec(`
		ALTER TABLE budgets
		ADD CONSTRAINT fk_budgets_user
		FOREIGN KEY (user_id) REFERENCES users(id)
		ON DELETE CASCADE;
	`)

	_, _ = DB.Exec(`
		ALTER TABLE budgets
		ADD CONSTRAINT fk_budgets_category
		FOREIGN KEY (category_id) REFERENCES categories(id)
		ON DELETE CASCADE;
	`)

	log.Println("Database migration completed with relations!")
}