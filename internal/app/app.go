package app

import (
	"fmt"

	"github.com/giicoo/maratWebSite/configs"
	http_v1 "github.com/giicoo/maratWebSite/internal/delivery/http/v1"
	mongo_db "github.com/giicoo/maratWebSite/internal/repository/mongodb"
	"github.com/giicoo/maratWebSite/internal/server"
	"github.com/giicoo/maratWebSite/internal/service"
	hashFunc "github.com/giicoo/maratWebSite/pkg/hash_password"
	_ "github.com/mattn/go-sqlite3"
)

func Run() error {
	// tools
	hash := hashFunc.NewHashTools()

	// config
	cfg, err := configs.GetConfig("./configs/config.json")
	if err != nil {
		return err
	}

	// insert dependencies
	repo := mongo_db.NewStore(cfg)
	services := service.NewServices(repo, hash, cfg)
	handler := http_v1.NewHandler(services, cfg)

	// init handlers
	r := handler.InitHandlers()

	// init db
	if err := repo.InitDB(); err != nil {
		return err
	}

	// init and start server
	srv := server.NewServer(fmt.Sprintf("%v:%v", cfg.HOST, cfg.PORT), r)

	if err := srv.Start(); err != nil {
		return err
	}
	return nil
}
