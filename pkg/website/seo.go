package website

import "rkb/pkg/config"

// SEO contains all metadata for proper search engine optimization.
type SEO struct {
	Title       string
	Description string
	Canonical   string
	NoIndex     bool
	Image       string
	Type        string // website, article
	Author      string
	PublishedAt string
	Breadcrumbs []Breadcrumb
}

type Breadcrumb struct {
	Name string
	URL  string
}

// SiteConfig holds site-wide configuration.
type SiteConfig struct {
	URL                       string
	Name                      string
	Description               string
	DefaultImage              string
	TwitterHandle             string
	GoogleSearchConsoleVerify string
	GoogleAnalyticsID         string
	GoogleTagManagerID        string
	Language                  string
	OrgLogo                   string
}

func LoadSiteConfig(cfg *config.Config) SiteConfig {
	gtmID := ""
	if cfg.GoogleTagManagerEnabled {
		gtmID = cfg.GoogleTagManagerID
	}
	return SiteConfig{
		URL:                       cfg.SiteURL,
		Name:                      cfg.SiteName,
		Description:               cfg.SiteDescription,
		DefaultImage:              "/static/og-image.png",
		TwitterHandle:             cfg.TwitterHandle,
		GoogleSearchConsoleVerify: cfg.GoogleSearchConsole,
		GoogleAnalyticsID:         cfg.GoogleAnalyticsID,
		GoogleTagManagerID:        gtmID,
		Language:                  "en",
		OrgLogo:                   "/static/logos/logo.svg",
	}
}
