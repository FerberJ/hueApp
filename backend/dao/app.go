package dao

import (
	"context"
	"openHueApp/backend/connection"
)

// App struct
type App struct {
	ctx     context.Context
	SqlLite *connection.SqliteDb
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	a.SqlLite = connection.NewSqliteDb("./mydb.db")
}
