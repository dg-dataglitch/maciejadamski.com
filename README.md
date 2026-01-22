# maciej-adamski

Personal website and blog built with **Go**, **templ**, and a minimal Tailwind setup (via CDN). Static pages are generated at build time and served from the `dist/` folder.

## What this repo contains

- ⚡ **Static output** for fast load times and easy hosting
- 📝 **Markdown blog** with frontmatter metadata
- 🔍 **SEO basics**: canonical URLs, Open Graph, Twitter cards, sitemap, robots.txt
- 🔁 **Dev server** with live rebuilds

## How it works

- `cmd/build` renders templates and Markdown into `dist/`.
- `cmd/dev` runs a local server and rebuilds on changes (via Air).
- `templates/` contains templ components; `content/posts/` contains Markdown posts.

## Quick start

```
# Clone
git clone https://github.com/<your-username>/<repo>.git
cd <repo>/website

# Install Go deps
go mod download

# Setup environment
cp .env.example .env

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
├── cmd/build/        # Build command
├── cmd/dev/          # Development server
├── content/posts/    # Markdown blog posts
├── templates/        # Templ HTML templates
│   ├── layouts/      # Base layout
│   ├── components/   # Navbar, Footer
│   ├── website/      # Homepage
│   └── blog/         # Blog templates
├── static/           # Static assets (CSS, icons, robots/sitemap templates)
├── dist/             # Generated output (gitignored)
└── .env              # Configuration
```

## Configuration

Edit `.env`:

```env
SITE_NAME=My Website
SITE_URL=https://example.com
SITE_DESCRIPTION=My personal website
GOOGLE_TAG_MANAGER_ENABLED=false
GOOGLE_ANALYTICS_ID=G-XXXXXXXXXX
```

## Writing blog posts

Create a markdown file in `content/posts/`:

```
---
title: "My Post Title"
description: "A brief description"
date: "2026-01-15T12:00:00Z"
published: true
---

Your content here...
```

## Customizing styles

Tailwind is included via CDN in [templates/layouts/base.templ](templates/layouts/base.templ). Update classes directly in templates or extend the inline Tailwind theme block.

## Deployment

Build the static site:

```
make build
```

Deploy the `dist/` folder to any static hosting:

- **GitHub Pages**
- **Cloudflare Pages**
- **Netlify**
- **Vercel**

## Requirements

- Go 1.21+
- Air (for hot reload): `go install github.com/air-verse/air@latest`
- Templ: `go install github.com/a-h/templ/cmd/templ@latest`

## License

MIT
