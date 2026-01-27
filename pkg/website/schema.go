package website

import (
	"encoding/json"
	"log/slog"
)

// OrganizationSchema generates JSON-LD for Organization (used on homepage).
func OrganizationSchema(site SiteConfig) string {
	org := map[string]any{
		"@context":    "https://schema.org",
		"@type":       "Organization",
		"name":        site.Name,
		"url":         site.URL,
		"description": site.Description,
	}
	if site.OrgLogo != "" {
		org["logo"] = map[string]any{
			"@type": "ImageObject",
			"url":   AbsoluteURL(site, site.OrgLogo),
		}
	}
	if site.OrgFounder != "" {
		org["founder"] = map[string]any{
			"@type": "Person",
			"name":  site.OrgFounder,
		}
	}
	if len(site.OrgSameAs) > 0 {
		org["sameAs"] = site.OrgSameAs
	}
	return marshalSchema(org)
}

// WebSiteSchema generates JSON-LD for WebSite (enables sitelinks in Google).
func WebSiteSchema(site SiteConfig) string {
	ws := map[string]any{
		"@context": "https://schema.org",
		"@type":    "WebSite",
		"name":     site.Name,
		"url":      site.URL,
	}
	return marshalSchema(ws)
}

// ArticleSchema generates JSON-LD for Article.
func ArticleSchema(site SiteConfig, seo SEO, currentPath string) string {
	return articleSchema(site, seo, currentPath, "Article")
}

// BlogPostingSchema generates JSON-LD for BlogPosting.
func BlogPostingSchema(site SiteConfig, seo SEO, currentPath string) string {
	return articleSchema(site, seo, currentPath, "BlogPosting")
}

// BreadcrumbSchema generates JSON-LD for BreadcrumbList.
func BreadcrumbSchema(site SiteConfig, breadcrumbs []Breadcrumb) string {
	if len(breadcrumbs) == 0 {
		return ""
	}
	items := make([]map[string]any, len(breadcrumbs))
	for i, bc := range breadcrumbs {
		items[i] = map[string]any{
			"@type":    "ListItem",
			"position": i + 1,
			"name":     bc.Name,
			"item":     AbsoluteURL(site, bc.URL),
		}
	}
	schema := map[string]any{
		"@context":        "https://schema.org",
		"@type":           "BreadcrumbList",
		"itemListElement": items,
	}
	return marshalSchema(schema)
}

// articleSchema is the shared implementation for Article and BlogPosting.
func articleSchema(site SiteConfig, seo SEO, currentPath string, schemaType string) string {
	article := map[string]any{
		"@context":    "https://schema.org",
		"@type":       schemaType,
		"headline":    seo.Title,
		"description": seo.Description,
		"url":         site.URL + currentPath,
		"image":       GetOGImage(site, seo),
		"mainEntityOfPage": map[string]any{
			"@type": "WebPage",
			"@id":   site.URL + currentPath,
		},
	}

	publisher := map[string]any{
		"@type": "Organization",
		"name":  site.Name,
	}
	if site.OrgLogo != "" {
		publisher["logo"] = map[string]any{
			"@type": "ImageObject",
			"url":   AbsoluteURL(site, site.OrgLogo),
		}
	}
	article["publisher"] = publisher

	if seo.ArticlePublishedTime != "" {
		article["datePublished"] = seo.ArticlePublishedTime
	}
	if seo.ArticleModifiedTime != "" {
		article["dateModified"] = seo.ArticleModifiedTime
	}
	if seo.ArticleAuthor != "" {
		author := map[string]any{
			"@type": "Person",
			"name":  seo.ArticleAuthor,
		}
		if seo.ArticleAuthorURL != "" {
			author["url"] = seo.ArticleAuthorURL
		}
		article["author"] = author
	}
	if seo.ArticleSection != "" {
		article["articleSection"] = seo.ArticleSection
	}
	if len(seo.ArticleTags) > 0 {
		article["keywords"] = seo.ArticleTags
	}
	return marshalSchema(article)
}

// marshalSchema converts a schema map to JSON string.
// Logs error and returns empty string on failure.
func marshalSchema(schema map[string]any) string {
	b, err := json.Marshal(schema)
	if err != nil {
		slog.Error("failed to marshal schema", "error", err)
		return ""
	}
	return string(b)
}
