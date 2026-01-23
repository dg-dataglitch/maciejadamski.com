package main

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"sort"
	"text/template"
	"time"

	"rkb/pkg/config"
	"rkb/pkg/generator"
	"rkb/pkg/logger"
	"rkb/pkg/markdown"
	"rkb/pkg/website"
	blogTemplates "rkb/templates/blog"
	webTemplates "rkb/templates/website"
)

func parsePostDate(date string) time.Time {
	if date == "" {
		return time.Time{}
	}
	parsed, err := time.Parse(time.RFC3339, date)
	if err == nil {
		return parsed
	}
	parsed, err = time.Parse("2006-01-02", date)
	if err == nil {
		return parsed
	}
	return time.Time{}
}

func latestPublishedPosts(posts []markdown.Post, limit int) []markdown.Post {
	if limit <= 0 {
		return []markdown.Post{}
	}
	sorted := make([]markdown.Post, 0, len(posts))
	for _, post := range posts {
		sorted = append(sorted, post)
	}
	sort.SliceStable(sorted, func(i, j int) bool {
		return parsePostDate(sorted[i].Meta.Date).After(parsePostDate(sorted[j].Meta.Date))
	})
	if len(sorted) > limit {
		sorted = sorted[:limit]
	}
	return sorted
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	l, err := logger.New(cfg.LogLevel)
	if err != nil {
		l, _ = logger.New("info")
	}
	slog.SetDefault(l)

	site := website.LoadSiteConfig(cfg)
	l.Debug("config_loaded", "log_level", cfg.LogLevel, "dist", cfg.DistPath, "static", cfg.StaticPath)
	l.Debug("site_config_loaded", "name", site.Name, "url", site.URL)
	l.Info("build_starting", "dist", cfg.DistPath)
	if err := os.RemoveAll(cfg.DistPath); err != nil {
		l.Error("cleanup_failed", "err", err)
	}
	if err := copyDir(cfg.StaticPath, filepath.Join(cfg.DistPath, "static")); err != nil {
		l.Error("static_copy_failed", "err", err)
	}
	postsDir := filepath.Join(cfg.ContentPath, "posts")
	posts, err := markdown.ParseDir(postsDir)
	if err != nil {
		l.Warn("blog_load_failed", "err", err)
		posts = []markdown.Post{}
	}
	var publishedPosts []markdown.Post
	for _, p := range posts {
		if p.Meta.Published {
			publishedPosts = append(publishedPosts, p)
		} else {
			l.Debug("post_skipped_unpublished", "slug", p.Slug)
		}
	}
	l.Debug("posts_filtered", "total", len(posts), "published", len(publishedPosts))
	sort.SliceStable(publishedPosts, func(i, j int) bool {
		return parsePostDate(publishedPosts[i].Meta.Date).After(parsePostDate(publishedPosts[j].Meta.Date))
	})
	latestPosts := latestPublishedPosts(publishedPosts, 3)
	homeSEO := website.SEO{
		Title:       site.Name + " - " + site.Description,
		Description: site.Description,
		Image:       site.DefaultImage,
	}
	l.Debug("rendering_home_page", "seo_title", homeSEO.Title)
	homePath := filepath.Join(cfg.DistPath, "index.html")
	if err := generator.RenderTemplComponent(homePath, webTemplates.Home(site, homeSEO, latestPosts)); err != nil {
		l.Error("home_render_failed", "err", err)
	} else {
		l.Info("page_generated", "path", "index.html")
	}
	indexSEO := website.SEO{
		Title:       "Blog | " + site.Name,
		Description: "Latest articles and insights from " + site.Name,
		Image:       site.DefaultImage,
	}
	indexPath := filepath.Join(cfg.DistPath, "blog", "index.html")
	if err := generator.RenderTemplComponent(indexPath, blogTemplates.Index(site, indexSEO, publishedPosts)); err != nil {
		l.Error("blog_index_failed", "err", err)
	} else {
		l.Info("page_generated", "path", "blog/index.html")
	}
	for _, post := range publishedPosts {
		postSEO := website.SEO{
			Title:       post.Meta.Title + " | " + site.Name,
			Description: post.Meta.Description,
			Type:        "article",
			PublishedAt: post.Meta.Date,
			Image:       site.DefaultImage,
		}
		postPath := filepath.Join(cfg.DistPath, "blog", post.Slug, "index.html")
		if err := generator.RenderTemplComponent(postPath, blogTemplates.PostPage(site, postSEO, post)); err != nil {
			l.Error("post_render_failed", "slug", post.Slug, "err", err)
		} else {
			l.Info("post_generated", "slug", post.Slug)
		}
	}
	generateSitemap(cfg.DistPath, cfg.StaticPath, site.URL, publishedPosts, l)
	generateRobots(cfg.DistPath, cfg.StaticPath, site.URL, l)
	l.Info("build_complete", "pages", 2, "posts", len(publishedPosts))
}

// copyDir recursively copies a directory tree, excluding source files.
func copyDir(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		if info.IsDir() && info.Name() == "scss" {
			return filepath.SkipDir
		}
		ext := filepath.Ext(path)
		if ext == ".tmpl" || ext == ".map" || ext == ".scss" {
			return nil
		}
		dstPath := filepath.Join(dst, relPath)
		if info.IsDir() {
			return os.MkdirAll(dstPath, 0755)
		}
		srcFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()
		if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err != nil {
			return err
		}
		dstFile, err := os.Create(dstPath)
		if err != nil {
			return err
		}
		defer dstFile.Close()
		_, err = io.Copy(dstFile, srcFile)
		return err
	})
}

func generateSitemap(distPath, staticPath, siteURL string, posts []markdown.Post, l *slog.Logger) {
	tmplPath := filepath.Join(staticPath, "sitemap.xml.tmpl")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		l.Error("sitemap_template_failed", "err", err)
		return
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
		l.Error("sitemap_create_failed", "err", err)
		return
	}
	defer f.Close()
	if err := tmpl.Execute(f, data); err != nil {
		l.Error("sitemap_execute_failed", "err", err)
	} else {
		l.Info("sitemap_generated")
	}
}

func generateRobots(distPath, staticPath, siteURL string, l *slog.Logger) {
	tmplPath := filepath.Join(staticPath, "robots.txt.tmpl")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		l.Error("robots_template_failed", "err", err)
		return
	}
	data := struct {
		SiteURL string
	}{
		SiteURL: siteURL,
	}
	outPath := filepath.Join(distPath, "robots.txt")
	f, err := os.Create(outPath)
	if err != nil {
		l.Error("robots_create_failed", "err", err)
		return
	}
	defer f.Close()
	if err := tmpl.Execute(f, data); err != nil {
		l.Error("robots_execute_failed", "err", err)
	} else {
		l.Info("robots_generated")
	}
}
