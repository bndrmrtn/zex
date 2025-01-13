package zex

import "net/http"

// Config is the configuration struct
type Config struct {
	// Development is a flag to enable development mode
	Development bool

	NotFoundHandler http.HandlerFunc
}

// make is a method to set the configuration
func defaultConfig() *Config {
	return &Config{
		Development: true,
	}
}

// make is a method to set the default configuration
func (c *Config) make() {
	if c.NotFoundHandler == nil {
		c.NotFoundHandler = http.NotFound
	}
}
