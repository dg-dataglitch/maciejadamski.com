package markdown

import (
	"bytes"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

// Post represents a parsed markdown file with metadata and rendered HTML content.
type Post struct {
	Meta    PostMeta
	Content string
}

// PostMeta contains frontmatter metadata from a markdown file.
type PostMeta struct {
	Title       string         `yaml:"title"`
	Date        string         `yaml:"date"`
	Description string         `yaml:"description"`
	Author      string         `yaml:"author"`
	Published   bool           `yaml:"published"`
	Slug        string         `yaml:"slug"`
	Extra       map[string]any `yaml:"-"`
}

// coreFields defines the standard frontmatter fields.
// Any field not in this set goes into Extra.
var coreFields = map[string]bool{
	"title":       true,
	"date":        true,
	"description": true,
	"author":      true,
	"published":   true,
	"slug":        true,
}

// FormattedDate formats the post date to a human-readable format.
// Supports RFC3339 and "2006-01-02" formats. Returns raw date on parse failure.
func (p PostMeta) FormattedDate() string {
	if p.Date == "" {
		return ""
	}
	t, err := time.Parse(time.RFC3339, p.Date)
	if err != nil {
		t, err = time.Parse("2006-01-02", p.Date)
		if err != nil {
			return p.Date
		}
	}
	return t.Format("January 02, 2006")
}

// ParseDate returns the parsed time.Time for sorting purposes.
// Returns zero time on parse failure.
func (p PostMeta) ParseDate() time.Time {
	t, err := time.Parse(time.RFC3339, p.Date)
	if err != nil {
		t, _ = time.Parse("2006-01-02", p.Date)
	}
	return t
}

// ParseDir reads all markdown files from a directory and returns parsed posts.
// Posts are sorted by date (newest first). Non-markdown files are ignored.
func ParseDir(dir string) ([]Post, error) {
	slog.Debug("parsing markdown directory", "dir", dir)

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("reading directory %s: %w", dir, err)
	}

	var posts []Post
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}
		path := filepath.Join(dir, entry.Name())
		post, err := ParseFile(path)
		if err != nil {
			slog.Warn("skipping file", "path", path, "error", err)
			continue
		}
		posts = append(posts, *post)
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Meta.ParseDate().After(posts[j].Meta.ParseDate())
	})

	slog.Debug("parsed markdown directory", "dir", dir, "count", len(posts))
	return posts, nil
}

// ParseFile reads a markdown file and extracts frontmatter and rendered HTML content.
func ParseFile(path string) (*Post, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	md := goldmark.New(goldmark.WithExtensions(meta.Meta))
	var buf bytes.Buffer
	ctx := parser.NewContext()

	if err := md.Convert(data, &buf, parser.WithContext(ctx)); err != nil {
		return nil, fmt.Errorf("parsing markdown: %w", err)
	}

	metaData := meta.Get(ctx)
	postMeta := extractMeta(metaData, path)

	slog.Debug("parsed file", "path", path, "title", postMeta.Title, "slug", postMeta.Slug)

	return &Post{Meta: postMeta, Content: buf.String()}, nil
}

// extractMeta converts raw metadata map to PostMeta struct.
func extractMeta(data map[string]any, path string) PostMeta {
	pm := PostMeta{
		Extra: make(map[string]any),
	}

	if v, ok := data["title"].(string); ok {
		pm.Title = v
	}
	if v, ok := data["description"].(string); ok {
		pm.Description = v
	}
	if v, ok := data["author"].(string); ok {
		pm.Author = v
	}
	if v, ok := data["published"].(bool); ok {
		pm.Published = v
	} else if v, ok := data["draft"].(bool); ok {
		pm.Published = !v
	}
	if v, ok := data["slug"].(string); ok {
		pm.Slug = v
	}

	// Handle date - can be string or time.Time depending on YAML parsing
	switch v := data["date"].(type) {
	case string:
		pm.Date = v
	case time.Time:
		pm.Date = v.Format("2006-01-02")
	}

	// Collect extra fields
	for key, value := range data {
		if !coreFields[key] {
			pm.Extra[key] = value
		}
	}

	// Generate slug from filename if not specified
	if pm.Slug == "" {
		pm.Slug = slugFromPath(path)
	}

	return pm
}

// slugFromPath generates a slug from a file path.
func slugFromPath(path string) string {
	filename := filepath.Base(path)
	return strings.TrimSuffix(filename, ".md")
}
