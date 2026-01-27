package website

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// ThemeConfig holds design tokens for the website.
type ThemeConfig struct {
	// Colors - UI
	ColorPageBackground string `yaml:"color_page_background"`
	ColorSurface        string `yaml:"color_surface"`
	ColorBorder         string `yaml:"color_border"`
	ColorBorderDark     string `yaml:"color_border_dark"`
	ColorLink           string `yaml:"color_link"`
	ColorBrand          string `yaml:"color_brand"`
	ColorScrollbar      string `yaml:"color_scrollbar"`

	// Colors - Typography
	ColorTextHeading string `yaml:"color_text_heading"`
	ColorTextBody    string `yaml:"color_text_body"`
	ColorTextMuted   string `yaml:"color_text_muted"`
	ColorTextBold    string `yaml:"color_text_bold"`
	ColorTextItalic  string `yaml:"color_text_italic"`
	ColorTextCode    string `yaml:"color_text_code"`
	ColorTextPre     string `yaml:"color_text_pre"`

	// Colors - Code Blocks
	ColorCodeBackground string `yaml:"color_code_background"`

	// Typography
	FontSans string `yaml:"font_sans"`
	FontMono string `yaml:"font_mono"`

	// Layout
	RadiusContainer string `yaml:"radius_container"`
}

// DefaultTheme returns sensible defaults.
func DefaultTheme() ThemeConfig {
	return ThemeConfig{
		ColorPageBackground: "#1c1c1d",
		ColorSurface:        "#343438",
		ColorBorder:         "#39393b",
		ColorBorderDark:     "#333333",
		ColorLink:           "#b2c3ff",
		ColorBrand:          "#52008d",
		ColorScrollbar:      "#555555",

		ColorTextHeading: "#e4e4e5",
		ColorTextBody:    "#909093",
		ColorTextMuted:   "#70757a",
		ColorTextBold:    "#ffffff",
		ColorTextItalic:  "#e4e4e5",
		ColorTextCode:    "#ffffff",
		ColorTextPre:     "#e4e4e5",

		ColorCodeBackground: "#212226",

		FontSans:        "'JetBrains Mono', monospace",
		FontMono:        "'JetBrains Mono', monospace",
		RadiusContainer: "12px",
	}
}

// LoadTheme loads theme configuration from a YAML file.
// Returns default theme if file doesn't exist.
func LoadTheme(path string) (ThemeConfig, error) {
	theme := DefaultTheme()

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return theme, nil
		}
		return theme, fmt.Errorf("reading theme file: %w", err)
	}

	if err := yaml.Unmarshal(data, &theme); err != nil {
		return theme, fmt.Errorf("parsing theme file: %w", err)
	}

	return theme, nil
}
