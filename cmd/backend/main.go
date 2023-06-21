package main

import (
	"context"
	"log"

	"github.com/spongeling/admin-api/internal/api"
	"github.com/spongeling/admin-api/internal/auth"
	"github.com/spongeling/admin-api/internal/freeling"
	"github.com/spongeling/admin-api/internal/gpt"
	"github.com/spongeling/admin-api/internal/repo"
	"github.com/spongeling/admin-api/internal/server"
	"github.com/spongeling/admin-api/internal/service"
	"github.com/spongeling/admin-api/shared"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// config
	err := shared.LoadConfig(".env")
	if err != nil {
		log.Fatalf("error loading config %v", err)
	}
	cfg, err := readConfig()
	if err != nil {
		log.Fatal(err)
	}

	// database
	db, err := repo.New(shared.GetDBConnectionString())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// auth
	users, err := db.GetAllUsers(ctx)
	if err != nil {
		log.Fatal(err)
	}
	auth.UpdateUsers(users)

	// 3rd party services
	gptClient := gpt.NewClient(ctx, cfg.GptApiKey)

	freelingClient, err := freeling.NewClient()
	if err != nil {
		log.Fatal("failed to connect to freeling server: ", err)
	}

	// service
	svc := service.New(db, gptClient, freelingClient)

	// apis
	pingApi := api.NewPing(svc)
	loginApi := api.NewLogin(svc)
	wordClassApi := api.NewWordClass(svc)
	categoryApi := api.NewCategory(svc)
	patternApi := api.NewPattern(svc)

	// server
	srv := server.New(cfg.HttpPort,
		pingApi,
		loginApi,
		categoryApi,
		wordClassApi,
		patternApi,
	)

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	err = srv.Start()
	if err != nil {
		log.Fatalf("start server err: %v", err)
	}
}
