package main

import (
	"database/sql"
	"log"

	auth "github.com/elecTRON-Fellowship/formula-1/internal/auth/api"
	db "github.com/elecTRON-Fellowship/formula-1/internal/auth/db/sqlc"
	config "github.com/elecTRON-Fellowship/formula-1/pkg/viper"
	_ "github.com/lib/pq"
)

func main() {
	config, err := config.LoadConfig("../../config")
	if err != nil {
		log.Fatal(err)
		return
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err)
		return
	}
	repo := db.NewRepo(conn)
	server, err := auth.NewServer(repo)
	if err != nil {
		log.Fatal(err)
		return
	}
	server.SetRoutes()
	if err = server.StartServer(config.Addr); err != nil {
		log.Fatal(err)
		return
	}
}
