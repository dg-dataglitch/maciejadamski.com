package generator

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/a-h/templ"
)

// RenderTemplComponent renders a templ component to an HTML file.
// Creates parent directories if they don't exist.
func RenderTemplComponent(path string, component templ.Component) error {
	slog.Debug("rendering component", "path", path)

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("creating directory %s: %w", dir, err)
	}

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("creating file %s: %w", path, err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			slog.Error("failed to close file", "path", path, "error", err)
		}
	}()

	if err := component.Render(context.Background(), f); err != nil {
		return fmt.Errorf("rendering component to %s: %w", path, err)
	}

	return nil
}
