# ADR-001: Markdown Package Core Design

Date: 2026-01-25
Status: Accepted
Related Decisions: None

## 1. Context

We are building gopress, a static site generator engine that will be used to create multiple website instances for different clients. The markdown package is a core component responsible for parsing blog posts from markdown files with YAML frontmatter.

## 2. Problem

The markdown package must be flexible enough to serve different clients with varying frontmatter needs (tags, categories, custom fields) while maintaining type safety for common fields and predictable behavior across all instances.

## 3. Decision

### D-001: Fixed Core Fields with Extra Map for Extensions

The `PostMeta` struct contains type-safe fields for common frontmatter (Title, Date, Description, Author, Published, Draft, Slug) plus an `Extra map[string]interface{}` field to capture client-specific fields without losing type safety on the core.

### D-002: Frontmatter Slug Overrides Filename

When determining the post slug: if `slug` is explicitly set in frontmatter, use it. Otherwise, derive the slug from the filename. This provides sensible defaults with an escape hatch for manual control.

### D-003: Safe Publishing Defaults

If `draft: true` is set, the post is unpublished regardless of the `published` field. If neither `draft` nor `published` is set, default to unpublished. Nothing goes live by accident.

## 4. Rationale

D-001 balances type safety (95% case) with flexibility (5% case). Fully dynamic maps lose compile-time checks. Hardcoding all possible fields violates YAGNI.

D-002 follows "convention over configuration". Most posts use filename-derived slugs. Explicit override handles edge cases like renamed files or SEO-optimized URLs.

D-003 prioritizes safety. Accidental publication is worse than accidental non-publication.

## 5. Consequences

- Client templates accessing custom fields must use type assertions on `Extra` map
- Parser implementation must populate `Extra` with all frontmatter keys not matching core fields
- All instances inherit the same publishing logic (no per-instance override)

## 6. Audit Log

- 2026-01-25: Document created and accepted.