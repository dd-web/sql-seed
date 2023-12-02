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

	defered_migrations = map[int][]byte{}

	result_account_verbose_prefix = "    - "
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

	fmt.Println("Seeding...")
	seeder := types.NewSeeder(store)
	seeder.Seed()

	fmt.Println("Finishing up...")
	finalize(store)

	// fmt.Printf("\nResults"+"\n-------\n"+"  * %v Accounts\n    * %v Admins\n    * %v Moderators\n    * %v Users\n", len(seeder.Accounts), len(seeder.Admins), len(seeder.Mods), len(seeder.Accounts)-(len(seeder.Admins)+len(seeder.Mods)))
	fmt.Printf(types.UnderlinePrint("Results")+"  * %v Accounts\n    * %v Admins\n    * %v Moderators\n    * %v Users\n", len(seeder.Accounts), len(seeder.Admins), len(seeder.Mods), len(seeder.Accounts)-(len(seeder.Admins)+len(seeder.Mods)))

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
		err := s.Execute(string(m.Up))
		if err != nil {
			log.Fatal(err)
		}

		m.Finished = true

		if m.Transatory != nil && len(m.Transatory) > 0 {
			defered_migrations[m.Index] = m.Transatory
			m.Finished = false
		}
	}

	fmt.Println("Migrations finished.")
}

// finalizes the migrations by running migrations defered by previous migrations
// these mostly consist of key constraints
func finalize(s *types.Store) {
	for _, m := range defered_migrations {
		err := s.Execute(string(m))
		if err != nil {
			fmt.Printf("Migration: %v\n\n", string(m))
			log.Fatal("Defered Migration Failure:", err.Error())
		}
	}
}
