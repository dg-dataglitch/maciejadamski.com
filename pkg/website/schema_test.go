package website

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestOrganizationSchema(t *testing.T) {
	site := SiteConfig{
		Name:        "Test Org",
		URL:         "https://example.com",
		Description: "A test organization",
		OrgLogo:     "/logo.png",
		OrgFounder:  "John Doe",
		OrgSameAs:   []string{"https://twitter.com/test"},
	}

	got := OrganizationSchema(site)

	// We check for key fields rather than exact JSON string match
	// to avoid brittleness with key ordering.
	if !strings.Contains(got, `"@type":"Organization"`) {
		t.Errorf("OrganizationSchema() missing @type Organization")
	}
	if !strings.Contains(got, `"name":"Test Org"`) {
		t.Errorf("OrganizationSchema() missing name")
	}
	if !strings.Contains(got, `"founder":{`) {
		t.Errorf("OrganizationSchema() missing founder")
	}
	if !strings.Contains(got, `"sameAs":["https://twitter.com/test"]`) {
		t.Errorf("OrganizationSchema() missing sameAs")
	}
}

func TestWebSiteSchema(t *testing.T) {
	site := SiteConfig{
		Name: "Test Site",
		URL:  "https://example.com",
	}

	got := WebSiteSchema(site)

	if !strings.Contains(got, `"@type":"WebSite"`) {
		t.Errorf("WebSiteSchema() missing @type WebSite")
	}
	if !strings.Contains(got, `"name":"Test Site"`) {
		t.Errorf("WebSiteSchema() missing name")
	}
}

func TestBreadcrumbSchema(t *testing.T) {
	site := SiteConfig{URL: "https://example.com"}
	crumbs := []Breadcrumb{
		{Name: "Home", URL: "/"},
		{Name: "Blog", URL: "/blog"},
	}

	got := BreadcrumbSchema(site, crumbs)

	if !strings.Contains(got, `"@type":"BreadcrumbList"`) {
		t.Errorf("BreadcrumbSchema() missing @type BreadcrumbList")
	}
	if !strings.Contains(got, `"position":1`) {
		t.Errorf("BreadcrumbSchema() missing position 1")
	}
	if !strings.Contains(got, `"position":2`) {
		t.Errorf("BreadcrumbSchema() missing position 2")
	}
}

func TestArticleSchema(t *testing.T) {
	site := SiteConfig{
		Name:    "Test Site",
		URL:     "https://example.com",
		OrgLogo: "/logo.png",
	}
	seo := SEO{
		Title:                "Test Article",
		Description:          "Description",
		ArticlePublishedTime: "2023-01-01",
		ArticleAuthor:        "Jane Doe",
	}

	got := ArticleSchema(site, seo, "/post")

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(got), &data); err != nil {
		t.Fatalf("ArticleSchema() produced invalid JSON: %v", err)
	}

	if data["@type"] != "Article" {
		t.Errorf("ArticleSchema() type = %v, want Article", data["@type"])
	}
	if data["headline"] != "Test Article" {
		t.Errorf("ArticleSchema() headline = %v, want Test Article", data["headline"])
	}
}

func TestBlogPostingSchema(t *testing.T) {
	site := SiteConfig{
		Name:    "Test Site",
		URL:     "https://example.com",
		OrgLogo: "/logo.png",
	}
	seo := SEO{
		Title: "Test Post",
	}

	got := BlogPostingSchema(site, seo, "/post")

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(got), &data); err != nil {
		t.Fatalf("BlogPostingSchema() produced invalid JSON: %v", err)
	}

	if data["@type"] != "BlogPosting" {
		t.Errorf("BlogPostingSchema() type = %v, want BlogPosting", data["@type"])
	}
}
