package generator

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCopyFile(t *testing.T) {
	tmpDir := t.TempDir()
	srcFile := filepath.Join(tmpDir, "source.txt")
	dstFile := filepath.Join(tmpDir, "dest.txt")

	content := []byte("hello world")
	if err := os.WriteFile(srcFile, content, 0644); err != nil {
		t.Fatal(err)
	}

	if err := CopyFile(srcFile, dstFile); err != nil {
		t.Fatalf("CopyFile() error = %v", err)
	}

	got, err := os.ReadFile(dstFile)
	if err != nil {
		t.Fatalf("ReadFile(dest) error = %v", err)
	}

	if string(got) != string(content) {
		t.Errorf("content = %q, want %q", string(got), string(content))
	}
}

func TestCopyDir(t *testing.T) {
	srcDir := t.TempDir()
	dstDir := t.TempDir()

	// Setup source structure
	// src/
	//   file1.txt
	//   subdir/
	//     file2.txt

	subDir := filepath.Join(srcDir, "subdir")
	if err := os.Mkdir(subDir, 0755); err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(filepath.Join(srcDir, "file1.txt"), []byte("file1"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(subDir, "file2.txt"), []byte("file2"), 0644); err != nil {
		t.Fatal(err)
	}

	// Copy to a new destination folder inside dstDir
	targetDir := filepath.Join(dstDir, "target")
	if err := CopyDir(srcDir, targetDir); err != nil {
		t.Fatalf("CopyDir() error = %v", err)
	}

	// Verify
	file1Path := filepath.Join(targetDir, "file1.txt")
	if _, err := os.Stat(file1Path); err != nil {
		t.Errorf("file1.txt missing: %v", err)
	}

	file2Path := filepath.Join(targetDir, "subdir", "file2.txt")
	if _, err := os.Stat(file2Path); err != nil {
		t.Errorf("subdir/file2.txt missing: %v", err)
	}
}
