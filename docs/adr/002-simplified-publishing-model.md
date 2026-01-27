# ADR-002: Simplified Publishing Model

Date: 2026-01-25
Status: Accepted
Related Decisions: Supersedes D-003

## 1. Context

During implementation of the markdown package per ADR-001, we identified that the `PostMeta` struct contained both `Draft` and `Published` fields. These are boolean inverses of each other, creating redundancy and potential for conflicting states.

## 2. Problem

Having both `Draft` and `Published` fields creates ambiguity:
- What happens when `draft: true` AND `published: true`?
- Extra logic required to resolve conflicts
- Two fields to maintain for the same concept
- Violates simplicity principle

## 3. Decision

### D-004: Single Published Field with Safe Default

Remove the `Draft` field entirely. Keep only `Published`. Go's zero value for bool is `false`, meaning posts are unpublished by default unless explicitly marked `published: true` in frontmatter.

The `PostMeta` struct becomes:
```go
type PostMeta struct {
    Title       string                 `yaml:"title"`
    Date        string                 `yaml:"date"`
    Description string                 `yaml:"description"`
    Author      string                 `yaml:"author"`
    Published   bool                   `yaml:"published"`
    Slug        string                 `yaml:"slug"`
    Extra       map[string]interface{} `yaml:"-"`
}
```

## 4. Rationale

- One field instead of two eliminates ambiguity
- Go zero value provides safe default (unpublished)
- No conflict resolution logic needed
- Positive semantics ("is this published?") clearer than negative ("is this NOT a draft?")
- YAGNI: we don't need Hugo/Jekyll compatibility

## 5. Consequences

- Existing content using `draft: true` will have that field land in `Extra` map (ignored for publishing logic)
- Migration path: replace `draft: true` with `published: false` (or just remove it, since unpublished is the default)
- Simpler mental model for content authors

## 6. Audit Log

- 2026-01-25: Document created, supersedes D-003 from ADR-001.