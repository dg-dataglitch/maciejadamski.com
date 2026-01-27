package engine

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"testing"

	"maciejadamski/pkg/markdown"
	"maciejadamski/pkg/website"

	"github.com/a-h/templ"
)

func TestDefaultOptions(t *testing.T) {
	opts := DefaultOptions()
	if opts.OutputDir != "dist" {
		t.Errorf("OutputDir = %v, want dist", opts.OutputDir)
	}
	if opts.ConfigDir != "config" {
		t.Errorf("ConfigDir = %v, want config", opts.ConfigDir)
	}
}

func TestFilterPublished(t *testing.T) {
	posts := []markdown.Post{
		{
			Meta: markdown.PostMeta{Title: "Published 1", Published: true},
		},
		{
			Meta: markdown.PostMeta{Title: "Draft 1", Published: false},
		},
		{
			Meta: markdown.PostMeta{Title: "Published 2", Published: true},
		},
	}

	got := filterPublished(posts)

	if len(got) != 2 {
		t.Errorf("len(got) = %d, want 2", len(got))
	}
	if got[0].Meta.Title != "Published 1" {
		t.Errorf("got[0].Title = %v, want Published 1", got[0].Meta.Title)
	}
	if got[1].Meta.Title != "Published 2" {
		t.Errorf("got[1].Title = %v, want Published 2", got[1].Meta.Title)
	}
}

// mockComponent implements templ.Component for testing
type mockComponent struct {
	content string
}

func (m mockComponent) Render(ctx context.Context, w io.Writer) error {
	_, err := w.Write([]byte(m.content))
	return err
}

func TestBuild_Basic(t *testing.T) {
	tmpDir := t.TempDir()
	configDir := filepath.Join(tmpDir, "config")
	contentDir := filepath.Join(tmpDir, "content")
	outputDir := filepath.Join(tmpDir, "dist")
	staticDir := filepath.Join(tmpDir, "static")

	// Setup directories
	os.MkdirAll(configDir, 0755)
	os.MkdirAll(filepath.Join(contentDir, "blog"), 0755)
	os.MkdirAll(staticDir, 0755)

	// Create dummy config
	os.WriteFile(filepath.Join(configDir, "site.yaml"), []byte("name: Test Site\nurl: http://test.com"), 0644)
	os.WriteFile(filepath.Join(configDir, "theme.yaml"), []byte("background: red"), 0644)

	// Create static file
	os.WriteFile(filepath.Join(staticDir, "style.css"), []byte("body { color: red; }"), 0644)

	opts := BuildOptions{
		OutputDir:  outputDir,
		ConfigDir:  configDir,
		ContentDir: contentDir,
		StaticDir:  staticDir,
	}

	components := ComponentRegistry{
		Index: func(c website.SiteConfig, s website.SEO, posts []markdown.Post) templ.Component {
			return mockComponent{content: "<h1>Home</h1>"}
		},
	}

	if err := Build(components, opts); err != nil {
		t.Fatalf("Build() error = %v", err)
	}

	// Verify output
	if _, err := os.Stat(filepath.Join(outputDir, "index.html")); err != nil {
		t.Error("index.html missing")
	}
	if _, err := os.Stat(filepath.Join(outputDir, "static", "style.css")); err != nil {
		t.Error("static/style.css missing")
	}
}

func TestBuild_MissingConfig(t *testing.T) {
	tmpDir := t.TempDir()
	opts := BuildOptions{
		OutputDir: filepath.Join(tmpDir, "dist"),
		ConfigDir: filepath.Join(tmpDir, "missing_config"),
	}

	components := ComponentRegistry{}

	err := Build(components, opts)
	if err == nil {
		t.Error("Build() expected error for missing config, got nil")
	}
}
