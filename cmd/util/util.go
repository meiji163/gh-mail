package util

import (
	"os"
	"path/filepath"
	"runtime"
)

func dataDir() string {
	var path string
	if a := os.Getenv("XDG_DATA_HOME"); a != "" {
		path = filepath.Join(a, "gh")
	} else if b := os.Getenv("LocalAppData"); runtime.GOOS == "windows" && b != "" {
		path = filepath.Join(b, "GitHub CLI")
	} else {
		c, _ := os.UserHomeDir()
		path = filepath.Join(c, ".local", "share", "gh")
	}
	return path
}

func configDir() string {
	var path string
	if a := os.Getenv("GH_CONFIG_DIR"); a != "" {
		path = a
	} else if b := os.Getenv("XDG_CONFIG_HOME"); b != "" {
		path = filepath.Join(b, "gh")
	} else if c := os.Getenv("AppData"); runtime.GOOS == "windows" && c != "" {
		path = filepath.Join(c, "GitHub CLI")
	} else {
		d, _ := os.UserHomeDir()
		path = filepath.Join(d, ".config", "gh")
	}
	return path
}

func ExtensionsDir() string {
	return filepath.Join(dataDir(), "extensions")
}

func PrivateKeyPath() string {
	return filepath.Join(dataDir(), "extensions", "gh-mail", "private.pem")
}
