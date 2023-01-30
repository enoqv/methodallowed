package MethodAllowed

import (
	"context"
	"net/http"
)

type Config struct {
	Methods []string `json:"Methods,omitempty" toml:"Methods,omitempty" yaml:"Methods,omitempty" export:"true"`
	Message string   `json:"Message,omitempty" toml:"Message,omitempty" yaml:"Message,omitempty" export:"true"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		Methods: []string{},
		Message: "",
	}
}

type MethodAllowed struct {
	cfg  *Config
	next http.Handler
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &MethodAllowed{
		cfg:  config,
		next: next,
	}, nil
}

func (m *MethodAllowed) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	for _, method := range m.cfg.Methods {
		if method == req.Method {
			m.next.ServeHTTP(rw, req)
			return
		}
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
	rw.Write([]byte(m.cfg.Message))
}
