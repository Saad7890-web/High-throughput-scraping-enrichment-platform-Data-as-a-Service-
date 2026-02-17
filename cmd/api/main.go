package main

import (
	"log"
	"net/http"

	"github.com/Saad7890-web/scrapper-platform/internal/config"
	"github.com/Saad7890-web/scrapper-platform/internal/db"
	"github.com/Saad7890-web/scrapper-platform/internal/job"
	"github.com/Saad7890-web/scrapper-platform/internal/server"
	"github.com/Saad7890-web/scrapper-platform/internal/worker"
)

func main(){
	cfg := config.Load()
	config.Must(cfg)
	dbConn := db.Connect(cfg.DBUrl)
	defer dbConn.Close()

	repo := job.NewRepository(dbConn)

	pool := worker.NewPool(10)

	executor := job.NewExecutor(repo)

	processor := worker.NewProcessor(executor)
	pool.Start(processor)

	service := job.NewService(repo, pool)
	handler := job.NewHandler(service)

	srv := server.New(handler)
	log.Println("server running on :" + cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, srv))
}