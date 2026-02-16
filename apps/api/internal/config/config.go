package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds application configuration from environment.
type Config struct {
	Domain     string
	Email      string
	DataDir    string
	Port       int
	JWTSecret  string
	AdminEmail string
}

// Load reads .env and builds Config from environment variables.
func Load() (*Config, error) {
	_ = godotenv.Load()

	port := 3001
	if p := os.Getenv("HELMOS_PORT"); p != "" {
		if v, err := strconv.Atoi(p); err == nil {
			port = v
		}
	}

	dataDir := "./data"
	if d := os.Getenv("HELMOS_DATA_DIR"); d != "" {
		dataDir = d
	}

	jwtSecret := "change-me-in-production"
	if s := os.Getenv("JWT_SECRET"); s != "" {
		jwtSecret = s
	}

	return &Config{
		Domain:     os.Getenv("HELMOS_DOMAIN"),
		Email:      os.Getenv("HELMOS_EMAIL"),
		DataDir:    dataDir,
		Port:       port,
		JWTSecret:  jwtSecret,
		AdminEmail: os.Getenv("ADMIN_EMAIL"),
	}, nil
}
