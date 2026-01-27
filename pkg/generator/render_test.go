package generator

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"testing"
)

type mockComponent struct {
	content string
}

func (m mockComponent) Render(ctx context.Context, w io.Writer) error {
	_, err := w.Write([]byte(m.content))
	return err
}

func TestRenderTemplComponent(t *testing.T) {
	tmpDir := t.TempDir()
	outFile := filepath.Join(tmpDir, "output.html")

	comp := mockComponent{content: "<html><body>Hello</body></html>"}

	if err := RenderTemplComponent(outFile, comp); err != nil {
		t.Fatalf("RenderTemplComponent() error = %v", err)
	}

	got, err := os.ReadFile(outFile)
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}

	if string(got) != comp.content {
		t.Errorf("content = %q, want %q", string(got), comp.content)
	}
}

func TestRenderTemplComponent_CreateDir(t *testing.T) {
	tmpDir := t.TempDir()
	// Output in a nested directory that doesn't exist
	outFile := filepath.Join(tmpDir, "nested", "dir", "output.html")

	comp := mockComponent{content: "nested"}

	if err := RenderTemplComponent(outFile, comp); err != nil {
		t.Fatalf("RenderTemplComponent() error = %v", err)
	}

	got, err := os.ReadFile(outFile)
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}

	if string(got) != "nested" {
		t.Errorf("content = %q, want nested", string(got))
	}
}
