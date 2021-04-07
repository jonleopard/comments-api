package main

import (
	"fmt"
	"net/http"

	"github.com/jonleopard/comments-api/internal/comment"
	"github.com/jonleopard/comments-api/internal/database"
	transportHTTP "github.com/jonleopard/comments-api/internal/transport/http"
)

// App - the struct which contains things link pointers
// to database connections
type App struct {
}

func main() {
	fmt.Println("Go REST API Course")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting up our APP")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	database.MigrateDB(db)
	if err != nil {
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to setup server")
		return err
	}

	return nil
}
