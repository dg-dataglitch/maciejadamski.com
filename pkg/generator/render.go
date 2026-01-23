package generator

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/a-h/templ"
)

func RenderTemplComponent(path string, component templ.Component) error {
	slog.Debug("rendering_component", "path", path)
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("creation directory failed %s: %w", dir, err)
	}
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("creation file failed %s: %w", path, err)
	}
	defer func(f *os.File) {
		cErr := f.Close()
		if err != nil && cErr != nil {
			slog.Error("file_close_failed", "path", path, "err", cErr)
		}
	}(f)
	return component.Render(context.Background(), f)
}
