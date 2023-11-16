package database

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Migration struct {
	Index    int
	Up       []byte
	Down     []byte
	Finished bool
}

// reads bytes from a file and returns them
func readbytes(path string) ([]byte, error) {
	bs, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

// creates the migration from the given path, populate the up and down fields and return it
func createMigration(path string, index int) (*Migration, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var migration Migration = Migration{Finished: false, Index: index}

	for _, item := range entries {
		if item.IsDir() {
			return nil, fmt.Errorf("found directory in migration directory: %s", item.Name())
		}

		bs, err := readbytes(path + "/" + item.Name())
		if err != nil {
			return nil, err
		}

		switch item.Name() {
		case "up.sql":
			migration.Up = bs
		case "down.sql":
			migration.Down = bs
		default:
			return nil, fmt.Errorf("found file in migration directory that is not up.sql or down.sql: %s", item.Name())
		}
	}
	return &migration, nil
}

// parse the directory for migration paths and create a map of migrations
func parseMigrationDir(path string) (map[int]*Migration, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var migrations map[int]*Migration = make(map[int]*Migration)

	for _, item := range entries {
		if item.IsDir() {
			match, err := regexp.Match("\\d{5}_.*", []byte(item.Name()))
			if err != nil {
				return nil, err
			}

			if match {
				iter, err := strconv.Atoi(item.Name()[0:5])
				if err != nil {
					return nil, err
				}

				migration, err := createMigration(path+"/"+item.Name(), iter)
				if err != nil {
					return nil, err
				}

				migrations[iter] = migration
			}
		}
	}

	return migrations, nil
}

// parse migrations directory and return a map of migrations that can be executed
// migrations are sorted by their index in ascending order, which is the order they should be executed in
// migrations include both the up and down sql scripts in bytes (should be converted to string before execution)
func Migrations(path string) (map[int]*Migration, error) {
	return parseMigrationDir(path)
}
