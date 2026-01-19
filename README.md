# go-ssg-template

A minimal, fast static site generator built with **Go**, **Templ**, and **Bootstrap 5.3**.

## Features

- ⚡ **Lightning Fast** - Pure static HTML, no client-side JS frameworks
- 🌙 **Dark/Light Mode** - Built-in theme switcher with Bootstrap 5.3 color modes
- 📱 **Mobile-First** - No hover effects or transitions for better mobile performance
- 🎨 **Bootstrap 5.3** - Modern styling with SCSS customization
- 📝 **Markdown Blog** - Write posts in markdown with frontmatter
- 🔍 **SEO Ready** - Meta tags, Open Graph, sitemap, robots.txt
- 🔥 **Hot Reload** - Air-powered development server

## Quick Start

```bash
# Clone
git clone https://github.com/YOUR_USERNAME/go-ssg-template.git my-site
cd my-site

# Install dependencies
go mod download
npm install

# Setup
cp .env.example .env

# Build CSS and site
make css
make build

# Development with hot reload
make dev
```

Open http://localhost:3000

### Testing on Mobile

When you run `make dev`, the server binds to `0.0.0.0` and shows your local network IP:

```
> Local:   http://localhost:3000
> Network: http://192.168.x.x:3000
```

Open the Network URL on your phone (same WiFi) to test mobile directly.

## Project Structure

```
├── cmd/build/        # Build command
├── cmd/dev/          # Development server
├── content/posts/    # Markdown blog posts
├── templates/        # Templ HTML templates
│   ├── layouts/      # Base layout
│   ├── components/   # Navbar, Footer
│   ├── website/      # Homepage
│   └── blog/         # Blog templates
├── static/           # Static assets
│   ├── scss/         # Bootstrap customization
│   └── css/          # Compiled CSS
├── dist/             # Generated output (gitignored)
└── .env              # Configuration
```

## Configuration

Edit `.env`:

```env
SITE_NAME=My Website
SITE_URL=https://example.com
SITE_DESCRIPTION=My awesome site
GOOGLE_ANALYTICS_ID=G-XXXXXXXXXX
```

## Writing Blog Posts

Create a markdown file in `content/posts/`:

```markdown
---
title: "My Post Title"
description: "A brief description"
date: "2024-01-15"
author: "Your Name"
published: true
---

Your content here...
```

## Customizing Styles

Edit `static/scss/main.scss` to customize Bootstrap variables:

```scss
// Colors
$primary: #000000;

// Typography
$font-family-base: 'Inter', system-ui, sans-serif;
```

Then rebuild CSS:

```bash
make css
```

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

- Go 1.21+
- Node.js 18+ (for SCSS compilation)
- Air (for hot reload): `go install github.com/air-verse/air@latest`
- Templ: `go install github.com/a-h/templ/cmd/templ@latest`

## License

MIT
