package repo

import (
	"errors"
	"os"
	"path/filepath"
)

const markerFile = "volon.yaml"

// FindRepoRoot walks up from start until it finds a directory containing volon.yaml.
func FindRepoRoot(start string) (string, error) {
	path := start
	for {
		if fileExists(filepath.Join(path, markerFile)) {
			return path, nil
		}

		parent := filepath.Dir(path)
		if parent == path {
			return "", errors.New("volon.yaml not found in this directory or any parent")
		}
		path = parent
	}
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}
