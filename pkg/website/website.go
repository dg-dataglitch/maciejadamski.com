package website

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// FontFamily represents supported Google Fonts.
type FontFamily string

const (
	FontGoogleSansFlex FontFamily = "google-sans-flex"
)

// Breadcrumb represents a single breadcrumb navigation item.
type Breadcrumb struct {
	Name string
	URL  string
}

// SiteConfig holds site-wide configuration loaded from YAML.
type SiteConfig struct {
	// Basic site info
	Name        string `yaml:"name"`
	URL         string `yaml:"url"`
	Description string `yaml:"description"`
	Language    string `yaml:"language"`
	Locale      string `yaml:"locale"`

	// Default images for SEO
	DefaultImage       string `yaml:"default_image"`
	DefaultImageAlt    string `yaml:"default_image_alt"`
	DefaultImageWidth  string `yaml:"default_image_width"`
	DefaultImageHeight string `yaml:"default_image_height"`

	// Branding
	OrgLogo    string   `yaml:"org_logo"`
	Favicon    string   `yaml:"favicon"`
	OrgFounder string   `yaml:"org_founder"`
	OrgSameAs  []string `yaml:"org_same_as"`

	// Social and analytics
	TwitterHandle             string `yaml:"twitter_handle"`
	GoogleSearchConsoleVerify string `yaml:"google_search_console_verify"`
	GoogleAnalyticsID         string `yaml:"google_analytics_id"`

	// Frontend features
	EnableAlpineJS bool       `yaml:"enable_alpine_js"`
	EnableHTMX     bool       `yaml:"enable_htmx"`
	FontFamily     FontFamily `yaml:"font_family"`
	CustomCSS      []string   `yaml:"custom_css"`

	// Theme (loaded separately, not from site.yaml)
	Theme ThemeConfig `yaml:"-"`
}

// SEO contains all metadata for rendering a single page.
type SEO struct {
	// Basic meta tags
	Title       string
	Description string
	Canonical   string
	NoIndex     bool

	// Open Graph
	OGType        string
	OGImage       string
	OGImageAlt    string
	OGImageWidth  string
	OGImageHeight string

	// Twitter Card
	TwitterCard    string
	TwitterCreator string

	// Article-specific (for blog posts)
	ArticlePublishedTime string
	ArticleModifiedTime  string
	ArticleAuthor        string
	ArticleAuthorURL     string
	ArticleSection       string
	ArticleTags          []string

	// Navigation
	Breadcrumbs []Breadcrumb

	// Page type flags (for conditional rendering in templates)
	IsHomePage  bool
	IsArticle   bool
	IsBlogIndex bool
}

// LoadSiteConfig reads site configuration from a YAML file.
func LoadSiteConfig(path string) (SiteConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return SiteConfig{}, fmt.Errorf("reading site config: %w", err)
	}
	var cfg SiteConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return SiteConfig{}, fmt.Errorf("parsing site config: %w", err)
	}
	return cfg, nil
}
