package main

import (
	"database/sql"
	"log"

	auth "github.com/elecTRON-Fellowship/formula-1/api"
	db "github.com/elecTRON-Fellowship/formula-1/database/sqlc"
	config "github.com/elecTRON-Fellowship/formula-1/pkg/viper"
	_ "github.com/lib/pq"
)

func main() {
	config, err := config.LoadConfig(".")
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
	server, err := auth.NewServer(config, repo)
	if err != nil {
		log.Fatal(err)
		return
	}
	if err := server.SetRoutes(); err != nil {
		log.Fatal(err)
	}
	if err = server.StartServer(config); err != nil {
		log.Fatal(err)
		return
	}
}
