package main

import (
	"fmt"
	"log"

	"github.com/dd-web/pgsvk-seeder/pkg/database"
	"github.com/dd-web/pgsvk-seeder/pkg/types"
)

var (
	pgHost = "david.local"
	pgUser = "postgres"
	pgPass = "admin"
	pgName = "local_db_test"
	pgPort = "5432"
	pgSSL  = "disable"

	migration_path = "./cmd/migrations"

	RUN_MIGRATIONS = false
)

func main() {
	fmt.Println("Starting...")
	config := types.NewPGConfig(pgUser, pgPass, pgName, pgHost, pgPort, pgSSL)
	store, err := types.NewStore(config)
	if err != nil {
		log.Fatal(err)
	}

	if RUN_MIGRATIONS {
		run_migrations(store)
	}
}

func run_migrations(s *types.Store) {
	migs, err := database.Migrations(migration_path)
	if err != nil {
		log.Fatal(err)
	}

	for _, migration := range migs {
		upStr := string(migration.Up)
		err := s.Execute(upStr)
		if err != nil {
			log.Fatal(err)
		}
		migration.Finished = true
	}
}
