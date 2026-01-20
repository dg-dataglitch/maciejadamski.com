package markdown

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

// PostMeta contains the frontmatter metadata for a blog post.
type PostMeta struct {
	Title       string
	Date        string
	Description string
	Author      string
	Published   bool
	Slug        string
}

// Post represents a parsed markdown blog post.
type Post struct {
	Meta    PostMeta
	Content string // HTML content
	Slug    string // URL-safe identifier
}

// FormattedDate returns the date in a human-readable format (e.g. "January 02, 2006").
// It attempts to parse the date as RFC3339 or "2006-01-02".
func (p Post) FormattedDate() string {
	t, err := time.Parse(time.RFC3339, p.Meta.Date)
	if err != nil {
		t, err = time.Parse("2006-01-02", p.Meta.Date)
		if err != nil {
			return p.Meta.Date // Return raw string if parsing fails
		}
	}
	return t.Format("January 02, 2006")
}

// ParseFile reads a markdown file and extracts frontmatter and content.
func ParseFile(path string) (*Post, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading file %s: %w", path, err)
	}

	md := goldmark.New(
		goldmark.WithExtensions(meta.Meta),
	)

	var buf bytes.Buffer
	ctx := parser.NewContext()

	if err := md.Convert(data, &buf, parser.WithContext(ctx)); err != nil {
		return nil, fmt.Errorf("parsing markdown %s: %w", path, err)
	}

	metaData := meta.Get(ctx)
	postMeta := extractMeta(metaData, path)

	return &Post{
		Meta:    postMeta,
		Content: buf.String(),
		Slug:    postMeta.Slug,
	}, nil
}

// ParseDir reads all markdown files from a directory.
func ParseDir(dir string) ([]Post, error) {
	var posts []Post

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("reading directory %s: %w", dir, err)
	}

	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".md" {
			continue
		}

		post, err := ParseFile(filepath.Join(dir, entry.Name()))
		if err != nil {
			continue // Skip files that can't be parsed
		}

		posts = append(posts, *post)
	}

	return posts, nil
}

// extractMeta converts raw metadata map to PostMeta struct.
func extractMeta(data map[string]interface{}, path string) PostMeta {
	pm := PostMeta{
		Published: false,
	}

	if title, ok := data["title"].(string); ok {
		pm.Title = title
	}
	if date, ok := data["date"].(string); ok {
		pm.Date = date
	}
	if desc, ok := data["description"].(string); ok {
		pm.Description = desc
	}
	if author, ok := data["author"].(string); ok {
		pm.Author = author
	}
	if published, ok := data["published"].(bool); ok {
		pm.Published = published
	} else if draft, ok := data["draft"].(bool); ok {
		// If published is not set, check for draft
		pm.Published = !draft
	} else {
		// Default to false if neither is set, or true?
		// Usually safely default to false (unpublished)
		pm.Published = false
	}
	if slug, ok := data["slug"].(string); ok {
		pm.Slug = slug
	} else {
		// Generate slug from filename if not specified
		base := filepath.Base(path)
		pm.Slug = strings.TrimSuffix(base, filepath.Ext(base))
	}

	return pm
}
