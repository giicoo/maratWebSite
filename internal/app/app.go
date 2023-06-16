package app

import (
	"database/sql"

	http_v1 "github.com/giicoo/maratWebSite/internal/delivery/http/v1"
	"github.com/giicoo/maratWebSite/internal/repository/sqlite"
	"github.com/giicoo/maratWebSite/internal/server"
	"github.com/giicoo/maratWebSite/internal/service"
	_ "github.com/mattn/go-sqlite3"
)

func Run() error {

	// open db
	db, err := sql.Open("sqlite3", "./repo.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// insert dependencies
	repo := sqlite.NewStore(db)
	services := service.NewServices(repo)
	handler := http_v1.NewHandler(services)

	// init handlers
	r := handler.InitHandlers()

	// init db
	if err := repo.InitDB(); err != nil {
		return err
	}

	// init and start server
	srv := server.NewServer("localhost:8000", r)

	if err := srv.Start(); err != nil {
		return err
	}
	return nil
}
