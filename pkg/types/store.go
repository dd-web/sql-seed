package types

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// postgres connection string format
// used to create a connection string for connecting to postgres
var sql_conn_fmt string = "user=%s password=%s dbname=%s host=%s port=%s sslmode=%s"

// debug mode for printing sql results
var debug_mode = false

// postgres connection config object
type PGConfig struct {
	User     string
	Password string
	Name     string
	Host     string
	Port     string
	SSL      string
}

// NewPGConfig creates a new PGConfig object
func NewPGConfig(user, pass, name, host, port, ssl string) *PGConfig {
	return &PGConfig{
		User:     user,
		Password: pass,
		Name:     name,
		Host:     host,
		Port:     port,
		SSL:      ssl,
	}
}

// ConnectionString returns a formatted string for connecting to postgres
func (pgc *PGConfig) ConnectionString() string {
	return fmt.Sprintf(sql_conn_fmt, pgc.User, pgc.Password, pgc.Name, pgc.Host, pgc.Port, pgc.SSL)
}

// Store holds a connection to the database
type Store struct {
	DB *sql.DB
}

// creates a new store object with a connection to the database
func NewStore(pgc *PGConfig) (*Store, error) {
	db, err := sql.Open("postgres", pgc.ConnectionString())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("could not ping database: %v", err)
	}

	return &Store{DB: db}, nil
}

// execute raw sql query
func (s *Store) Execute(query string) error {
	res, err := s.DB.Exec(query)
	if err != nil {
		return err
	}

	if debug_mode {
		fmt.Printf("[SQL RESULT]: %+v\n", res)
	}
	return nil
}

type SQLStringer interface {
	Insert() string
	// Select() string
	// Update() string
	// Delete() string
}

func GetSQLInsert(s SQLStringer) string {
	return s.Insert()
}

// func GetSQLSelect(s SQLStringer) string {
// 	return s.Select()
// }

func (s *Store) SeedThread(account_id int) error {
	identity := &Identity{
		AccountID: account_id,
		RoleID:    3,                // 3 for creator
		StyleID:   1,                // 1 for default
		Name:      "X7DmfZF9ddzggj", // random string for alias
	}
	str := identity.Insert()
	// create identity for account

	// create thread_content

	// create thread

	// create thread_identity

	return nil
}
