package engine

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"maciejadamski/pkg/generator"
	"maciejadamski/pkg/markdown"
	"maciejadamski/pkg/website"

	"github.com/a-h/templ"
)

// ComponentRegistry holds the templ components provided by the user's site.
// The engine calls these functions to render pages.
type ComponentRegistry struct {
	// Index renders the homepage.
	Index func(website.SiteConfig, website.SEO, []markdown.Post) templ.Component

	// BlogIndex renders the blog listing page (optional).
	BlogIndex func(website.SiteConfig, website.SEO, []markdown.Post) templ.Component

	// BlogPost renders individual blog posts.
	BlogPost func(website.SiteConfig, website.SEO, markdown.Post) templ.Component
}

// BuildOptions configures the build process paths.
type BuildOptions struct {
	OutputDir  string
	ConfigDir  string
	ContentDir string
	StaticDir  string
}

// DefaultOptions returns standard build paths.
func DefaultOptions() BuildOptions {
	return BuildOptions{
		OutputDir:  "dist",
		ConfigDir:  "config",
		ContentDir: "content",
		StaticDir:  "static",
	}
}

// Build generates the static site using the provided components and options.
func Build(components ComponentRegistry, opts BuildOptions) error {
	slog.Info("starting build", "output", opts.OutputDir)

	if err := os.RemoveAll(opts.OutputDir); err != nil {
		return fmt.Errorf("cleaning output directory: %w", err)
	}

	if err := os.MkdirAll(opts.OutputDir, 0755); err != nil {
		return fmt.Errorf("creating output directory: %w", err)
	}

	site, err := loadSiteWithTheme(opts.ConfigDir)
	if err != nil {
		return err
	}

	blogDir := filepath.Join(opts.ContentDir, "blog")
	posts, err := markdown.ParseDir(blogDir)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("parsing blog directory: %w", err)
	}
	publishedPosts := filterPublished(posts)

	if components.Index != nil {
		seo := website.SEO{
			Title:       site.Name,
			Description: site.Description,
			IsHomePage:  true,
		}
		outputPath := filepath.Join(opts.OutputDir, "index.html")
		slog.Debug("rendering homepage", "path", outputPath)
		if err := generator.RenderTemplComponent(outputPath, components.Index(site, seo, publishedPosts)); err != nil {
			return fmt.Errorf("rendering homepage: %w", err)
		}
	}

	if err := buildBlog(components, opts, site, publishedPosts); err != nil {
		return err
	}

	if err := copyStaticFiles(opts); err != nil {
		return err
	}

	if err := GenerateSitemap(opts.OutputDir, opts.StaticDir, site.URL, publishedPosts); err != nil {
		slog.Warn("failed to generate sitemap", "error", err)
	}

	if err := GenerateRobots(opts.OutputDir, opts.StaticDir, site.URL); err != nil {
		slog.Warn("failed to generate robots.txt", "error", err)
	}

	slog.Info("build completed", "output", opts.OutputDir)
	return nil
}

// loadSiteWithTheme loads site config and theme from the config directory.
func loadSiteWithTheme(configDir string) (website.SiteConfig, error) {
	site, err := website.LoadSiteConfig(filepath.Join(configDir, "site.yaml"))
	if err != nil {
		return website.SiteConfig{}, fmt.Errorf("loading site config: %w", err)
	}

	theme, err := website.LoadTheme(filepath.Join(configDir, "theme.yaml"))
	if err != nil {
		slog.Warn("theme load failed, using defaults", "error", err)
		theme = website.DefaultTheme()
	}
	site.Theme = theme

	return site, nil
}

// buildBlog renders the blog index and all published blog posts.
func buildBlog(components ComponentRegistry, opts BuildOptions, site website.SiteConfig, published []markdown.Post) error {
	if len(published) == 0 {
		return nil
	}

	if components.BlogIndex != nil {
		seo := website.SEO{
			Title:       "Blog - " + site.Name,
			Description: site.Description,
			IsBlogIndex: true,
		}
		indexPath := filepath.Join(opts.OutputDir, "blog", "index.html")

		slog.Debug("rendering blog index", "path", indexPath, "posts", len(published))

		if err := generator.RenderTemplComponent(indexPath, components.BlogIndex(site, seo, published)); err != nil {
			return fmt.Errorf("rendering blog index: %w", err)
		}
	}

	if components.BlogPost == nil {
		slog.Warn("no BlogPost component provided, skipping post rendering")
		return nil
	}

	for _, post := range published {
		seo := website.SEO{
			Title:                post.Meta.Title + " - " + site.Name,
			Description:          post.Meta.Description,
			IsArticle:            true,
			ArticlePublishedTime: post.Meta.Date,
			ArticleAuthor:        post.Meta.Author,
		}
		postPath := filepath.Join(opts.OutputDir, "blog", post.Meta.Slug, "index.html")

		slog.Debug("rendering blog post", "slug", post.Meta.Slug, "path", postPath)

		if err := generator.RenderTemplComponent(postPath, components.BlogPost(site, seo, post)); err != nil {
			return fmt.Errorf("rendering blog post %s: %w", post.Meta.Slug, err)
		}
	}

	slog.Info("blog built", "posts", len(published))
	return nil
}

// filterPublished returns only posts with Published=true.
func filterPublished(posts []markdown.Post) []markdown.Post {
	var published []markdown.Post
	for _, p := range posts {
		if p.Meta.Published {
			published = append(published, p)
		}
	}
	return published
}

// copyStaticFiles copies the static directory to the output.
func copyStaticFiles(opts BuildOptions) error {
	src := opts.StaticDir
	dst := filepath.Join(opts.OutputDir, "static")

	if _, err := os.Stat(src); os.IsNotExist(err) {
		slog.Debug("no static directory found", "path", src)
		return nil
	}

	slog.Debug("copying static files", "from", src, "to", dst)

	if err := generator.CopyDir(src, dst); err != nil {
		return fmt.Errorf("copying static files: %w", err)
	}

	return nil
}
