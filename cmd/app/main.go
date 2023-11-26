package main

import (
	"fmt"
	"log"

	"github.com/dd-web/pgsvk-seeder/pkg/database"
	"github.com/dd-web/pgsvk-seeder/pkg/types"
)

var (
	migration_enable   = true
	migration_path     = "./cmd/migrations"
	migration_rollback = true
)

func main() {
	fmt.Println("Starting...")

	store, err := types.NewStore()
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
