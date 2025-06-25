package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config is the root configuration structure for the entire application.
type Config struct {
	Server      ServerConfig
	Database    DatabaseConfig
	ExternalAPI ExternalAPIConfig
}

func (c *Config) Validate() error {
	if c.Server.Port == "" {
		return fmt.Errorf("server port cannot be empty")
	}
	if c.Database.Host == "" || c.Database.User == "" || c.Database.DBName == "" {
		return fmt.Errorf("database host, user, or dbname cannot be empty")
	}
	if c.Database.Port == 0 {
		return fmt.Errorf("database port cannot be zero")
	}
	if c.ExternalAPI.OpenBeautyFactsBaseURL == "" {
		return fmt.Errorf("OpenBeautyFacts base URL cannot be empty")
	}

	return nil
}

// LoadConfig loads configuration from environment variables, applying defaults.
func LoadConfig() (*Config, error) {
	cfg := &Config{}

	// --- Server Configuration ---
	cfg.Server.Port = getEnvOrDefault("SERVER_PORT", DefaultServerPort)

	// --- Database Configuration ---
	cfg.Database.Host = getEnvOrDefault("SUPABASE_HOST", DefaultDBHost)
	cfg.Database.User = getEnvOrDefault("SUPABASE_USER", DefaultDBUser)
	cfg.Database.Password = os.Getenv("SUPABASE_PASSWORD")
	cfg.Database.DBName = getEnvOrDefault("SUPABASE_DB_NAME", DefaultDBName)

	dbPortStr := getEnvOrDefault("SUPABASE_PORT", fmt.Sprintf("%d", DefaultDBPort))
	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		return nil, fmt.Errorf("invalid SUPABASE_PORT '%s': %w", dbPortStr, err)
	}
	cfg.Database.Port = dbPort

	// --- External API Configuration ---
	cfg.ExternalAPI.OpenBeautyFactsBaseURL = getEnvOrDefault("OPEN_BEAUTY_FACTS_BASE_URL", DefaultOpenBeautyFactsBaseURL)

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("configuration validation failed: %w", err)
	}

	return cfg, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
