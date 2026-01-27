package website

import "strings"

// AbsoluteURL joins the site URL with a path, ensuring a valid absolute URL.
func AbsoluteURL(site SiteConfig, path string) string {
	if path == "" {
		return ""
	}
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		return path
	}
	base := strings.TrimRight(site.URL, "/")
	cleanPath := strings.TrimLeft(path, "/")
	return base + "/" + cleanPath
}

// GetLanguage returns the site language or defaults to "en".
func GetLanguage(site SiteConfig) string {
	if site.Language != "" {
		return site.Language
	}
	return "en"
}

// GetLocale returns the site locale or defaults to "en_US".
func GetLocale(site SiteConfig) string {
	if site.Locale != "" {
		return site.Locale
	}
	return "en_US"
}

// GetRobots returns the robots meta tag content.
func GetRobots(seo SEO) string {
	if seo.NoIndex {
		return "noindex, nofollow"
	}
	return "index, follow"
}

// GetCanonical returns the canonical URL for the page.
func GetCanonical(site SiteConfig, seo SEO, currentPath string) string {
	if seo.Canonical != "" {
		return seo.Canonical
	}
	return AbsoluteURL(site, currentPath)
}

// GetOGType returns the Open Graph type or defaults to "website".
func GetOGType(seo SEO) string {
	if seo.OGType != "" {
		return seo.OGType
	}
	return "website"
}

// GetOGImage determines the best Open Graph image to use.
func GetOGImage(site SiteConfig, seo SEO) string {
	if seo.OGImage != "" {
		return AbsoluteURL(site, seo.OGImage)
	}
	if site.DefaultImage != "" {
		return AbsoluteURL(site, site.DefaultImage)
	}
	return ""
}

// GetOGImageWidth returns the OG image width.
func GetOGImageWidth(site SiteConfig, seo SEO) string {
	if seo.OGImageWidth != "" {
		return seo.OGImageWidth
	}
	return site.DefaultImageWidth
}

// GetOGImageHeight returns the OG image height.
func GetOGImageHeight(site SiteConfig, seo SEO) string {
	if seo.OGImageHeight != "" {
		return seo.OGImageHeight
	}
	return site.DefaultImageHeight
}

// GetOGImageAlt returns the OG image alt text.
func GetOGImageAlt(site SiteConfig, seo SEO) string {
	if seo.OGImageAlt != "" {
		return seo.OGImageAlt
	}
	return site.DefaultImageAlt
}

// GetTwitterCard returns the Twitter card type or defaults to "summary_large_image".
func GetTwitterCard(seo SEO) string {
	if seo.TwitterCard != "" {
		return seo.TwitterCard
	}
	return "summary_large_image"
}

// GetTwitterSite returns the Twitter handle for the site.
func GetTwitterSite(site SiteConfig) string {
	return site.TwitterHandle
}

// GetTwitterCreator returns the Twitter handle for the content creator.
func GetTwitterCreator(site SiteConfig, seo SEO) string {
	if seo.TwitterCreator != "" {
		return seo.TwitterCreator
	}
	return site.TwitterHandle
}
