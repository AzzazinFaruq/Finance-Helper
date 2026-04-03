package Setup

import (
	"database/sql"
	"fmt"

	"simple_crud/Config"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"

	_ "github.com/go-sql-driver/mysql"
)

var DB *bun.DB

func Connect() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local",
		config.Get("DB_USER"),
		config.Get("DB_PASSWORD"),
		config.Get("DB_HOST"),
		config.Get("DB_PORT"),
		config.Get("DB_NAME"),
	)

	sqldb, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	DB = bun.NewDB(sqldb, mysqldialect.New())

	if err := sqldb.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("✅ MySQL connected")
}