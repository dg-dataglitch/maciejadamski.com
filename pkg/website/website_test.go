package website

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadSiteConfig(t *testing.T) {
	dir := t.TempDir()
	configFile := filepath.Join(dir, "site.yaml")

	content := []byte(`
name: My Awesome Site
url: https://mysite.com
language: en
`)
	if err := os.WriteFile(configFile, content, 0644); err != nil {
		t.Fatal(err)
	}

	cfg, err := LoadSiteConfig(configFile)
	if err != nil {
		t.Fatalf("LoadSiteConfig() error = %v", err)
	}

	if cfg.Name != "My Awesome Site" {
		t.Errorf("cfg.Name = %v, want My Awesome Site", cfg.Name)
	}
	if cfg.URL != "https://mysite.com" {
		t.Errorf("cfg.URL = %v, want https://mysite.com", cfg.URL)
	}
}

func TestLoadSiteConfig_NotFound(t *testing.T) {
	_, err := LoadSiteConfig("/non/existent/site.yaml")
	if err == nil {
		t.Error("LoadSiteConfig() expected error for missing file, got nil")
	}
}

func TestLoadSiteConfig_InvalidYAML(t *testing.T) {
	dir := t.TempDir()
	configFile := filepath.Join(dir, "site.yaml")

	// Invalid YAML content (tab indentation is not allowed in YAML, but let's use something definitely broken)
	content := []byte(`
name: My Site
  url: broken indentation
`)
	if err := os.WriteFile(configFile, content, 0644); err != nil {
		t.Fatal(err)
	}

	_, err := LoadSiteConfig(configFile)
	if err == nil {
		t.Error("LoadSiteConfig() expected error for invalid YAML, got nil")
	}
}
