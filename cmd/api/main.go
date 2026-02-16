package main

import (
	"log"
	"net/http"

	"github.com/Saad7890-web/scrapper-platform/internal/config"
	"github.com/Saad7890-web/scrapper-platform/internal/db"
	"github.com/Saad7890-web/scrapper-platform/internal/server"
)

func main(){
	cfg := config.Load()
	config.Must(cfg)
	dbConn := db.Connect(cfg.DBUrl)
	defer dbConn.Close()

	srv := server.New()
	log.Println("server running on :" + cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, srv))
}