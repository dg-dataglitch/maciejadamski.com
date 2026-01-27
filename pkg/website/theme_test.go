package website

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDefaultTheme(t *testing.T) {
	theme := DefaultTheme()
	if theme.ColorPageBackground != "#1c1c1d" {
		t.Errorf("expected default ColorPageBackground #1c1c1d, got %s", theme.ColorPageBackground)
	}
	if theme.ColorTextHeading != "#e4e4e5" {
		t.Errorf("expected default ColorTextHeading #e4e4e5, got %s", theme.ColorTextHeading)
	}
	if theme.RadiusContainer != "12px" {
		t.Errorf("expected default RadiusContainer 12px, got %s", theme.RadiusContainer)
	}
}

func TestLoadTheme(t *testing.T) {
	dir := t.TempDir()
	themePath := filepath.Join(dir, "theme.yaml")

	// Create a dummy theme file
	content := []byte(`
color_page_background: "#000000"
color_text_heading: "#ff0000"
radius_container: "8px"
`)
	if err := os.WriteFile(themePath, content, 0644); err != nil {
		t.Fatal(err)
	}

	theme, err := LoadTheme(themePath)
	if err != nil {
		t.Fatalf("LoadTheme failed: %v", err)
	}

	if theme.ColorPageBackground != "#000000" {
		t.Errorf("expected loaded ColorPageBackground #000000, got %s", theme.ColorPageBackground)
	}
	if theme.ColorTextHeading != "#ff0000" {
		t.Errorf("expected loaded ColorTextHeading #ff0000, got %s", theme.ColorTextHeading)
	}
	if theme.RadiusContainer != "8px" {
		t.Errorf("expected loaded RadiusContainer 8px, got %s", theme.RadiusContainer)
	}
}

func TestLoadTheme_NotFound(t *testing.T) {
	// Should return default theme if file not found
	theme, err := LoadTheme("non_existent_theme.yaml")
	if err != nil {
		t.Fatalf("LoadTheme failed for missing file: %v", err)
	}

	defaults := DefaultTheme()
	if theme.ColorPageBackground != defaults.ColorPageBackground {
		t.Errorf("expected default theme when file missing")
	}
}
