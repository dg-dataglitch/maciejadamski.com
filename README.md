# maciej-adamski

Personal website and blog built with **Go**, **templ**, and **Tailwind CSS**. Static pages are generated at build time and served from the `dist/` folder.

## What this repo contains

- ⚡ **Static output** for fast load times and easy hosting
- 📝 **Markdown blog** with frontmatter metadata
- 🔍 **SEO basics**: canonical URLs, Open Graph, Twitter cards, sitemap, robots.txt
- 🎨 **Theming engine**: customizable colors and fonts via YAML
- REPEAT **Dev server** with live rebuilds (via Air)

## How it works

- `cmd/build`: renders templates and Markdown into `dist/` based on `config/`.
- `cmd/dev`: runs a local server and rebuilds on changes.
- `pkg/engine`: core logic for site generation, caching, and theming.
- `config/`: central location for site metadata and theme settings.

## Quick start

```bash
# Clone
git clone https://github.com/dg-dataglitch/maciejadamski.com.git
cd maciejadamski.com

# Install Go deps
go mod download

# Install git hooks (required for auto-building dist on commit)
make setup

# Run dev server (auto rebuild)
make dev
```

Open <http://localhost:3000>

### Testing on mobile

When you run `make dev`, the server binds to `0.0.0.0` and shows your local network IP:

```
> Local:   http://localhost:3000
> Network: http://192.168.x.x:3000
```

Open the Network URL on your phone (same WiFi) to test mobile directly.

## Project structure

```
├── cmd/
│   ├── build/        # Build command entry point
│   └── dev/          # Development server entry point
├── config/           # Site configuration (content & theme)
├── content/
│   └── posts/        # Markdown blog posts
├── pkg/              # Core logic
│   ├── engine/       # Build engine (site generation)
│   ├── markdown/     # Markdown parsing and processing
│   └── website/      # Site structs and models
├── templates/        # Templ HTML templates
│   ├── layouts/      # Base, HTML structure
│   ├── components/   # UI components (Navbar, Footer, etc.)
│   ├── pages/        # Page templates
│   └── blog/         # Blog-specific templates
├── static/           # Static assets (CSS, icons, images)
├── dist/             # Generated output (committed to repo)
└── Makefile          # Project commands
```

## Configuration

Site configuration is managed via YAML files in the `config/` directory.

### Site Metadata (`config/site.yaml`)

Define general site information, SEO defaults, and integrations.

```yaml
name: "My Website"
url: "https://example.com"
description: "My personal website"
language: "en"
twitter_handle: "@handle"
google_analytics_id: "G-XXXXXXXXXX"
```

### Theme (`config/theme.yaml`)

Customize the look and feel using CSS variables mapped to Tailwind colors.

```yaml
font_sans: "Inter, sans-serif"
font_mono: "JetBrains Mono, monospace"

# Colors (HSL format preferred for Tailwind compatibility)
color_brand: "217 91% 60%"  # Blue
color_surface: "220 14% 96%" # Light gray
# ... see file for full list
```

## Writing blog posts

Create a markdown file in `content/blog/`:

```markdown
---
title: "My Post Title"
description: "A brief description"
date: "2026-01-15T12:00:00Z"
published: true
author: "Me"
slug: "my-post-title"
---

Your content here...
```

## Customizing styles

- **Theme**: Edit `config/theme.yaml` to change global colors and fonts.
- **Custom CSS**: Edit `static/css/prose.css` (or other files in static) and reference them in `config/site.yaml`.
- **Tailwind**: Used via CDN in `templates/layouts/base.templ`. You can use arbitrary Tailwind classes in your templates.

## Deployment

Build the static site:

```bash
make build
```

Deploy the `dist/` folder to any static hosting:

- **GitHub Pages**
- **Cloudflare Pages**
- **Netlify**
- **Vercel**

## Requirements

- Go 1.22+
- Air (for hot reload): `go install github.com/air-verse/air@latest`
- Templ: `go install github.com/a-h/templ/cmd/templ@latest`

## License

MIT
