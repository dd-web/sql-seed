package types

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var seeder_debug_enabled = false

/******************************/
/* POSTGRES CONFIG / DEFAULTS */
/******************************/

var (
	pg_default_connfmt string = "user=%s password=%s dbname=%s host=%s port=%s sslmode=%s"
	pg_default_user    string = "postgres"
	pg_default_pass    string = "admin"
	pg_default_dbname  string = "local_db_test"
	pg_default_host    string = "david.local"
	pg_default_port    string = "5432"
	pg_default_ssl     string = "disable"
)

type pgConfigFunc func(*pgConfig) *pgConfig

type pgConfig struct {
	connfmt  string
	user     string
	password string
	name     string
	host     string
	port     string
	ssl      string
}

func defaultPGConfig() *pgConfig {
	return &pgConfig{
		connfmt:  pg_default_connfmt,
		user:     pg_default_user,
		password: pg_default_pass,
		name:     pg_default_dbname,
		host:     pg_default_host,
		port:     pg_default_port,
		ssl:      pg_default_ssl,
	}
}

/*******************************/
/*** CONFIGURATION FUNCTIONS ***/
/*******************************/

func PGCfgSetUser(s string) pgConfigFunc {
	return func(c *pgConfig) *pgConfig {
		c.user = s
		return c
	}
}

func PGCfgSetPass(s string) pgConfigFunc {
	return func(c *pgConfig) *pgConfig {
		c.password = s
		return c
	}
}

func PGCfgSetDBName(s string) pgConfigFunc {
	return func(c *pgConfig) *pgConfig {
		c.name = s
		return c
	}
}

func PGCfgSetHost(s string) pgConfigFunc {
	return func(c *pgConfig) *pgConfig {
		c.host = s
		return c
	}
}

func PGCfgSetPort(s string) pgConfigFunc {
	return func(c *pgConfig) *pgConfig {
		c.port = s
		return c
	}
}

func PGCfgSetSSL(s string) pgConfigFunc {
	return func(c *pgConfig) *pgConfig {
		c.ssl = s
		return c
	}
}

func PGCfgSetConnFmt(s string) pgConfigFunc {
	return func(c *pgConfig) *pgConfig {
		c.connfmt = s
		return c
	}
}

func newPGConfig(cfg ...pgConfigFunc) *pgConfig {
	config := defaultPGConfig()
	for _, fn := range cfg {
		config = fn(config)
	}
	return config
}

func (pgc *pgConfig) connstr() string {
	return fmt.Sprintf(pgc.connfmt, pgc.user, pgc.password, pgc.name, pgc.host, pgc.port, pgc.ssl)
}

/*********/
/* STORE */
/*********/

type Store struct {
	DB  *sql.DB
	cfg *pgConfig
}

func NewStore(cfg ...pgConfigFunc) (*Store, error) {
	pgcfg := newPGConfig(cfg...)
	db, err := sql.Open("postgres", pgcfg.connstr())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("could not ping database: %v", err)
	}

	return &Store{
		DB:  db,
		cfg: pgcfg,
	}, nil
}

func (s *Store) Execute(query string) error {
	res, err := s.DB.Exec(query)
	if err != nil {
		return err
	}

	if seeder_debug_enabled {
		fmt.Printf("[SQL RESULT]: %+v\n", res)
	}
	return nil
}
