package markdown

import "testing"

func TestPostMeta_FormattedDate(t *testing.T) {
	tests := []struct {
		name     string
		date     string
		expected string
	}{
		{
			name:     "RFC3339 format",
			date:     "2026-01-15T10:30:00Z",
			expected: "January 15, 2026",
		},
		{
			name:     "simple date format",
			date:     "2026-01-15",
			expected: "January 15, 2026",
		},
		{
			name:     "invalid date returns raw",
			date:     "not-a-date",
			expected: "not-a-date",
		},
		{
			name:     "empty date returns empty",
			date:     "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pm := PostMeta{Date: tt.date}
			got := pm.FormattedDate()
			if got != tt.expected {
				t.Errorf("FormattedDate() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestSlugFromPath(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "simple filename",
			path:     "my-post.md",
			expected: "my-post",
		},
		{
			name:     "path with directory",
			path:     "/content/posts/my-post.md",
			expected: "my-post",
		},
		{
			name:     "path with multiple directories",
			path:     "/home/user/blog/content/posts/another-post.md",
			expected: "another-post",
		},
		{
			name:     "filename with dots",
			path:     "my.special.post.md",
			expected: "my.special.post",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slugFromPath(tt.path)
			if got != tt.expected {
				t.Errorf("slugFromPath(%q) = %q, want %q", tt.path, got, tt.expected)
			}
		})
	}
}

func TestParseFile(t *testing.T) {
	tests := []struct {
		name      string
		file      string
		wantErr   bool
		checkPost func(t *testing.T, post *Post)
	}{
		{
			name:    "valid post with all fields",
			file:    "testdata/valid_post.md",
			wantErr: false,
			checkPost: func(t *testing.T, post *Post) {
				if post.Meta.Title != "Test Post Title" {
					t.Errorf("Title = %q, want %q", post.Meta.Title, "Test Post Title")
				}
				if post.Meta.Slug != "custom-slug" {
					t.Errorf("Slug = %q, want %q", post.Meta.Slug, "custom-slug")
				}
				if post.Meta.Published != true {
					t.Errorf("Published = %v, want true", post.Meta.Published)
				}
				if post.Meta.Author != "Test Author" {
					t.Errorf("Author = %q, want %q", post.Meta.Author, "Test Author")
				}
				if post.Content == "" {
					t.Error("Content should not be empty")
				}
			},
		},
		{
			name:    "minimal post uses filename as slug",
			file:    "testdata/minimal_post.md",
			wantErr: false,
			checkPost: func(t *testing.T, post *Post) {
				if post.Meta.Title != "Minimal Post" {
					t.Errorf("Title = %q, want %q", post.Meta.Title, "Minimal Post")
				}
				if post.Meta.Slug != "minimal_post" {
					t.Errorf("Slug = %q, want %q", post.Meta.Slug, "minimal_post")
				}
				if post.Meta.Published != false {
					t.Errorf("Published = %v, want false (default)", post.Meta.Published)
				}
			},
		},
		{
			name:    "post with extra fields",
			file:    "testdata/post_with_extras.md",
			wantErr: false,
			checkPost: func(t *testing.T, post *Post) {
				if len(post.Meta.Extra) != 2 {
					t.Errorf("Extra fields count = %d, want 2", len(post.Meta.Extra))
				}
				if _, ok := post.Meta.Extra["tags"]; !ok {
					t.Error("Extra should contain 'tags'")
				}
				if _, ok := post.Meta.Extra["category"]; !ok {
					t.Error("Extra should contain 'category'")
				}
			},
		},
		{
			name:      "non-existent file returns error",
			file:      "testdata/does_not_exist.md",
			wantErr:   true,
			checkPost: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			post, err := ParseFile(tt.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.checkPost != nil && post != nil {
				tt.checkPost(t, post)
			}
		})
	}
}

func TestParseDir(t *testing.T) {
	tests := []struct {
		name      string
		dir       string
		wantErr   bool
		wantCount int
	}{
		{
			name:      "valid directory with posts",
			dir:       "testdata",
			wantErr:   false,
			wantCount: 4,
		},
		{
			name:      "non-existent directory returns error",
			dir:       "testdata/does_not_exist",
			wantErr:   true,
			wantCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			posts, err := ParseDir(tt.dir)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(posts) != tt.wantCount {
				t.Errorf("ParseDir() returned %d posts, want %d", len(posts), tt.wantCount)
			}
		})
	}
}
