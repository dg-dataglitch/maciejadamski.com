package generator

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// CopyDir recursively copies a directory from src to dst.
// Symlinks are skipped. File permissions are preserved.
func CopyDir(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("stat source %s: %w", src, err)
	}

	if err := os.MkdirAll(dst, srcInfo.Mode()); err != nil {
		return fmt.Errorf("create destination %s: %w", dst, err)
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return fmt.Errorf("read directory %s: %w", src, err)
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if err := CopyDir(srcPath, dstPath); err != nil {
				return err
			}
			continue
		}

		// Skip symlinks
		if entry.Type()&fs.ModeSymlink != 0 {
			continue
		}

		if err := CopyFile(srcPath, dstPath); err != nil {
			return err
		}
	}

	return nil
}

// CopyFile copies a single file from src to dst.
// File permissions are preserved.
func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("open source %s: %w", src, err)
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("create destination %s: %w", dst, err)
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return fmt.Errorf("copy %s to %s: %w", src, dst, err)
	}

	info, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("stat source %s: %w", src, err)
	}

	if err := os.Chmod(dst, info.Mode()); err != nil {
		return fmt.Errorf("chmod destination %s: %w", dst, err)
	}

	return nil
}
