package database

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const driver = "postgres"

func DBConnect(dbName, user, password string) (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf(
		"user=%s dbname=%s password=%s sslmode=disable",
		user,
		dbName,
		password)
	return sqlx.Connect(driver, dataSourceName)
}

func CreateSchema(db *sqlx.DB, schema string) sql.Result {
	result := db.MustExec(schema)
	return result
}

func DropTable(db *sqlx.DB, tableName string) sql.Result {
	query := fmt.Sprintf("DROP TABLE IF EXISTS %s", tableName)
	return db.MustExec(query)
}
