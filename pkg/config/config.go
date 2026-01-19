package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	// Server
	Host     string
	Port     string
	LogLevel string

	// Paths
	DistPath    string
	ContentPath string
	StaticPath  string

	// Site
	SiteName        string
	SiteURL         string
	SiteDescription string

	// SEO & Analytics
	TwitterHandle       string
	GoogleSearchConsole string
	GoogleAnalyticsID   string
}

func Load() (*Config, error) {
	_ = godotenv.Load()
	cfg := &Config{
		Host:     getEnv("HOST", "localhost"),
		Port:     getEnv("PORT", "3000"),
		LogLevel: getEnv("LOG_LEVEL", "info"),

		DistPath:    getEnv("DIST_PATH", "dist"),
		ContentPath: getEnv("CONTENT_PATH", "content"),
		StaticPath:  getEnv("STATIC_PATH", "static"),

		SiteName:        getEnv("SITE_NAME", "My Website"),
		SiteURL:         getEnv("SITE_URL", "https://example.com"),
		SiteDescription: getEnv("SITE_DESCRIPTION", "A modern static website"),

		TwitterHandle:       getEnv("TWITTER_HANDLE", ""),
		GoogleSearchConsole: getEnv("GOOGLE_SEARCH_CONSOLE", ""),
		GoogleAnalyticsID:   getEnv("GOOGLE_ANALYTICS_ID", ""),
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) Addr() string {
	return c.Host + ":" + c.Port
}

func (c *Config) validate() error {
	if c.DistPath == "" {
		return fmt.Errorf("DIST_PATH cannot be empty")
	}
	return nil
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
