package commands

import "github.com/pkprzekwas/gotodo/database"

var id int

var dbName = "postgres"
var dbUser = "postgres"
var dbPass = "secret_pass"

func getTodoService() *database.TodoService {
	db, err := database.DBConnect(dbName, dbUser, dbPass)
	if err != nil {
		panic(err)
	}
	return database.CreateTodoService(db)
}
