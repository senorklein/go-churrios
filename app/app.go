package app

import (
	"database/sql"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)


type App struct {
	Router *mux.Router
	Database *sql.Database
}


func (app *App) SetupRouter() {

	#defines each route in the application
	app.Router.
			Methods("GET").
			Path("/user/{uuid}").
			HandlerFunc(app.getUserByUUID)
}

