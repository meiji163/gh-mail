package util

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

const defaultHost = "github.com"

type Config interface {
	Host() string
	Login() string
}

type config map[string]string

func (cfg config) Host() string {
	if host := os.Getenv("GH_HOST"); host != "" {
		return host
	}
	if host, ok := cfg["host"]; ok {
		return host
	}
	return defaultHost
}

func (cfg config) Login() string {
	return cfg["user"]
}

func NewConfig() Config {
	cfg := config(map[string]string{})
	hostMap := make(map[string]config)

	hostFile := filepath.Join(configDir(), "hosts.yml")
	hostData, err := os.ReadFile(hostFile)
	if err != nil {
		return cfg
	}
	if err := yaml.Unmarshal(hostData, &hostMap); err != nil {
		return cfg
	}

	for host, c := range hostMap {
		cfg["host"] = host
		cfg["user"] = c["user"]
	}
	return cfg
}
