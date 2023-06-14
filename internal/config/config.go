package config

type Config struct {
	Environment string
	Database    DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Port     uint
	DBName   string
	Username string
	Password string
}
