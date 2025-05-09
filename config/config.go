package config

import (
	"errors"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Algo     string    `yaml:"algo"`
	Host     string    `yaml:"host"`
	Port     string    `yaml:"port"`
	Server   Server    `yaml:"server"`
	Backends []Backend `yaml:"backends"`
}

type Server struct {
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:write_timeout"`
	IdleTimeout  time.Duration `yaml:"idle_timeout"`
}

type Backend struct {
	Url string `yaml:"url"`
}

func ReadConfigFile(path string) (*Config, error) {

	configFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil

}

func (cfg *Config) ValidateConfig() error {

	if len(cfg.Backends) == 0 {
		return errors.New("Cannot have zero backends, must have at least one")
	}
	if cfg.Algo == "" {
		cfg.Algo = "round-robin"
	}
	if cfg.Host == "" {
		cfg.Host = "localhost"
	}
	if cfg.Port == "" {
		cfg.Port = "8001"
	}

	err := cfg.Server.ConfigureServer()
	if err != nil {
		return err
	}

	for i := 0; i < len(cfg.Backends); i++ {
		backend := &cfg.Backends[i]
		backend.ConfigureBackend()
	}

	return nil
}

func (srvr *Server) ConfigureServer() error {

	// Read, Write and Idle timeouts will be defaulted to unlimited, so if they are not set
	// in the config then we leave the default. The unlimited default comes from the
	// Server in fasthttp package https://pkg.go.dev/github.com/valyala/fasthttp#Server

	// TODO: Add more fields for configuration of the Server
	return nil
}

func (b *Backend) ConfigureBackend() {
	// TODO: Add mopre fields for configuration such as Max # of connections, durations etc
}
