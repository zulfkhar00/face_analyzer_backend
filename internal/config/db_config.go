package config

// DatabaseConfig holds the configuration for the PostgreSQL database.
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string // e.g., "require", "disable", "prefer"
}
