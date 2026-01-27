package engine

import (
	"log/slog"
	"os"
	"path/filepath"
	"text/template"

	"maciejadamski/pkg/markdown"
)

// GenerateSitemap generates a sitemap.xml from a template in the static directory.
func GenerateSitemap(distPath, staticPath, siteURL string, posts []markdown.Post) error {
	tmplPath := filepath.Join(staticPath, "sitemap.xml.tmpl")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}

	data := struct {
		SiteURL string
		Posts   []markdown.Post
	}{
		SiteURL: siteURL,
		Posts:   posts,
	}

	outPath := filepath.Join(distPath, "sitemap.xml")
	f, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := tmpl.Execute(f, data); err != nil {
		return err
	}

	slog.Info("sitemap generated", "path", outPath)
	return nil
}

// GenerateRobots generates a robots.txt from a template in the static directory.
func GenerateRobots(distPath, staticPath, siteURL string) error {
	tmplPath := filepath.Join(staticPath, "robots.txt.tmpl")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}

	data := struct {
		SiteURL string
	}{
		SiteURL: siteURL,
	}

	outPath := filepath.Join(distPath, "robots.txt")
	f, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := tmpl.Execute(f, data); err != nil {
		return err
	}

	slog.Info("robots.txt generated", "path", outPath)
	return nil
}
