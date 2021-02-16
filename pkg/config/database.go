package config

import "fmt"

type DBConfig struct {
	host          string
	port          int
	name          string
	user          string
	password      string
	migrationPath string
}

func (db *DBConfig) Address() string {
	// user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", db.user, db.password,db.host, db.port, db.name)
}

func (db *DBConfig) MigrationPath() string {
	return db.migrationPath
}

func newDBConfig() DBConfig {
	return DBConfig{
		host:          getString("DB_HOST"),
		port:          getInt("DB_PORT"),
		name:          getString("DB_NAME"),
		user:          getString("DB_USER"),
		password:      getString("DB_PASSWORD"),
		migrationPath: getString("MIGRATION_PATH"),
	}
}
