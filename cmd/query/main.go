package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"github.com/JouleJoestar/web-10/internal/query/api"
	"github.com/JouleJoestar/web-10/internal/query/config"
	"github.com/JouleJoestar/web-10/internal/query/provider"
	"github.com/JouleJoestar/web-10/internal/query/usecase"
	_ "github.com/lib/pq"
)

func main() {
	configPath := flag.String("config-path", "config.yaml", "путь к файлу конфигурации")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dp := provider.NewProvider(db)
	uc := usecase.NewUsecase(dp)
	srv := api.NewServer(cfg.IP, cfg.Port, uc)

	fmt.Printf("Сервер запущен на %s\n", cfg.IP)
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
