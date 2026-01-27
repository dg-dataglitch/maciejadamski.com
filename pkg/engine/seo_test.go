package engine

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"maciejadamski/pkg/markdown"
)

func TestGenerateRobots(t *testing.T) {
	// Setup temporary directory structure
	tmpDir := t.TempDir()
	distDir := filepath.Join(tmpDir, "dist")
	staticDir := filepath.Join(tmpDir, "static")
	if err := os.MkdirAll(distDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(staticDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Create mock template
	tmplContent := "User-agent: *\nSitemap: {{ .SiteURL }}/sitemap.xml"
	if err := os.WriteFile(filepath.Join(staticDir, "robots.txt.tmpl"), []byte(tmplContent), 0644); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		siteURL string
		wantErr bool
		want    string
	}{
		{
			name:    "valid generation",
			siteURL: "https://example.com",
			wantErr: false,
			want:    "User-agent: *\nSitemap: https://example.com/sitemap.xml",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := GenerateRobots(distDir, staticDir, tt.siteURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateRobots() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				content, err := os.ReadFile(filepath.Join(distDir, "robots.txt"))
				if err != nil {
					t.Fatal(err)
				}
				if string(content) != tt.want {
					t.Errorf("file content = %q, want %q", string(content), tt.want)
				}
			}
		})
	}
}

func TestGenerateRobots_MissingTemplate(t *testing.T) {
	tmpDir := t.TempDir()
	err := GenerateRobots(tmpDir, tmpDir, "https://example.com")
	if err == nil {
		t.Error("GenerateRobots() expected error for missing template, got nil")
	}
}

func TestGenerateSitemap(t *testing.T) {
	tmpDir := t.TempDir()
	distDir := filepath.Join(tmpDir, "dist")
	staticDir := filepath.Join(tmpDir, "static")
	os.MkdirAll(distDir, 0755)
	os.MkdirAll(staticDir, 0755)

	// Mock template
	tmplContent := `<?xml version="1.0" encoding="UTF-8"?>
<urlset>
{{ range .Posts }}
  <url><loc>{{ $.SiteURL }}/blog/{{ .Meta.Slug }}</loc></url>
{{ end }}
</urlset>`
	os.WriteFile(filepath.Join(staticDir, "sitemap.xml.tmpl"), []byte(tmplContent), 0644)

	posts := []markdown.Post{
		{Meta: markdown.PostMeta{Slug: "post-1"}},
		{Meta: markdown.PostMeta{Slug: "post-2"}},
	}

	err := GenerateSitemap(distDir, staticDir, "https://example.com", posts)
	if err != nil {
		t.Fatalf("GenerateSitemap() error = %v", err)
	}

	content, err := os.ReadFile(filepath.Join(distDir, "sitemap.xml"))
	if err != nil {
		t.Fatal(err)
	}

	sContent := string(content)
	if !strings.Contains(sContent, "https://example.com/blog/post-1") {
		t.Error("sitemap missing post-1 URL")
	}
	if !strings.Contains(sContent, "https://example.com/blog/post-2") {
		t.Error("sitemap missing post-2 URL")
	}
}
