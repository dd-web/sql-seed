package main

import (
	"fmt"
	"log"

	"github.com/dd-web/pgsvk-seeder/pkg/database"
	"github.com/dd-web/pgsvk-seeder/pkg/types"
)

var (
	/* Postgres Configurations */

	// host of the postgres database - this can be an ip address or a domain name
	// mine is "david.local" because of how i have my local network setup. yours will
	// probably be "localhost" or "127.0.0.1" or something like that
	pgHost = "david.local"

	// username and password for the postgres database
	pgUser = "postgres"
	pgPass = "admin"

	// name of the database to use
	pgName = "local_db_test"

	// port and ssl settings
	pgPort = "5432"
	pgSSL  = "disable"

	/* Migration Configurations */

	// determines whether or not migrations will be ran
	migration_enable = true

	// path to migrations relative to the root of the project
	// each migration should contain an up and down sql file, inside the same directory
	// the containing folder should be prefixed with it's sequence number, e.g. 00001_create_table
	migration_path = "./cmd/migrations"

	// if true migrations will first be rolled back to nothing before being ran
	// this is useful for quickly resetting the database to a clean state and prototyping
	migration_rollback = true
)

func main() {
	fmt.Println("Starting...")
	config := types.NewPGConfig(pgUser, pgPass, pgName, pgHost, pgPort, pgSSL)
	store, err := types.NewStore(config)
	if err != nil {
		log.Fatal(err)
	}

	if migration_enable {
		migrate(store)
	}

	// SEED
	fmt.Println("Seeding...")
	err = store.SeedThread(1)
	if err != nil {
		log.Fatal(err)
	}

}

// runs migrations according to the configurations set above at the top of this file
func migrate(s *types.Store) {
	migrations, err := database.Migrations(migration_path)
	if err != nil {
		log.Fatal(err)
	}

	// if rollback, run down migrations in reverse order to reset the database to a clean state
	if migration_rollback {
		fmt.Println("Rolling back migrations...")
		for i := range migrations {
			m := migrations[len(migrations)-i+1]
			err := s.Execute(string(m.Down))
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	// run up migrations in sequence
	fmt.Println("Running migrations...")
	for _, m := range migrations {
		upStr := string(m.Up)
		err := s.Execute(upStr)
		if err != nil {
			log.Fatal(err)
		}
		m.Finished = true
	}

	fmt.Println("Migrations finished.")
}
