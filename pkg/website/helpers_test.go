package website

import (
	"testing"
)

func TestAbsoluteURL(t *testing.T) {
	site := SiteConfig{
		URL: "https://example.com",
	}

	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "empty path",
			path: "",
			want: "",
		},
		{
			name: "already absolute http",
			path: "http://other.com/foo",
			want: "http://other.com/foo",
		},
		{
			name: "already absolute https",
			path: "https://other.com/foo",
			want: "https://other.com/foo",
		},
		{
			name: "relative path",
			path: "blog/post",
			want: "https://example.com/blog/post",
		},
		{
			name: "relative path with leading slash",
			path: "/blog/post",
			want: "https://example.com/blog/post",
		},
		{
			name: "site url with trailing slash",
			path: "blog",
			want: "https://example.com/blog",
		},
	}

	// Test case for site URL with trailing slash
	siteTrailing := SiteConfig{URL: "https://example.com/"}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AbsoluteURL(site, tt.path)
			if got != tt.want {
				t.Errorf("AbsoluteURL() = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("site url with trailing slash", func(t *testing.T) {
		got := AbsoluteURL(siteTrailing, "blog")
		want := "https://example.com/blog"
		if got != want {
			t.Errorf("AbsoluteURL() = %v, want %v", got, want)
		}
	})
}

func TestGetLanguage(t *testing.T) {
	tests := []struct {
		name string
		site SiteConfig
		want string
	}{
		{
			name: "configured language",
			site: SiteConfig{Language: "fr"},
			want: "fr",
		},
		{
			name: "default language",
			site: SiteConfig{},
			want: "en",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLanguage(tt.site); got != tt.want {
				t.Errorf("GetLanguage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLocale(t *testing.T) {
	tests := []struct {
		name string
		site SiteConfig
		want string
	}{
		{
			name: "configured locale",
			site: SiteConfig{Locale: "fr_FR"},
			want: "fr_FR",
		},
		{
			name: "default locale",
			site: SiteConfig{},
			want: "en_US",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLocale(tt.site); got != tt.want {
				t.Errorf("GetLocale() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRobots(t *testing.T) {
	tests := []struct {
		name string
		seo  SEO
		want string
	}{
		{
			name: "no index",
			seo:  SEO{NoIndex: true},
			want: "noindex, nofollow",
		},
		{
			name: "index default",
			seo:  SEO{NoIndex: false},
			want: "index, follow",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRobots(tt.seo); got != tt.want {
				t.Errorf("GetRobots() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCanonical(t *testing.T) {
	site := SiteConfig{URL: "https://example.com"}
	tests := []struct {
		name string
		seo  SEO
		path string
		want string
	}{
		{
			name: "custom canonical",
			seo:  SEO{Canonical: "https://other.com/page"},
			path: "/page",
			want: "https://other.com/page",
		},
		{
			name: "generated canonical",
			seo:  SEO{},
			path: "/page",
			want: "https://example.com/page",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCanonical(site, tt.seo, tt.path); got != tt.want {
				t.Errorf("GetCanonical() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOGType(t *testing.T) {
	tests := []struct {
		name string
		seo  SEO
		want string
	}{
		{
			name: "custom type",
			seo:  SEO{OGType: "article"},
			want: "article",
		},
		{
			name: "default type",
			seo:  SEO{},
			want: "website",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOGType(tt.seo); got != tt.want {
				t.Errorf("GetOGType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOGImage(t *testing.T) {
	site := SiteConfig{
		URL:          "https://example.com",
		DefaultImage: "/default.jpg",
	}

	tests := []struct {
		name string
		seo  SEO
		want string
	}{
		{
			name: "seo image",
			seo:  SEO{OGImage: "/custom.jpg"},
			want: "https://example.com/custom.jpg",
		},
		{
			name: "default image",
			seo:  SEO{},
			want: "https://example.com/default.jpg",
		},
		{
			name: "no image at all",
			seo:  SEO{}, // site also has no default in this sub-test case logic?
			// Wait, I am using the 'site' variable which has DefaultImage.
			// I need to override site config for the "no image" case if I want to test it.
			want: "https://example.com/default.jpg",
		},
	}

	// Special case for no default image
	t.Run("no image configured", func(t *testing.T) {
		siteNoImg := SiteConfig{URL: "https://example.com"}
		seo := SEO{}
		if got := GetOGImage(siteNoImg, seo); got != "" {
			t.Errorf("GetOGImage() = %v, want empty string", got)
		}
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOGImage(site, tt.seo); got != tt.want {
				t.Errorf("GetOGImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
