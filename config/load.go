package config

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

// Application represents the configuration for the application
type Application struct {
	Server Server `toml:"server"`
}

// Server represents the configuration for running the MUD server
type Server struct {
	BindAddress string `toml:"bind_address"`
	BindPort    int    `toml:"bind_port"`
}

// Load loads the configuration from config.toml into an Application instance
func Load() Application {
	var file io.Reader
	var data []byte
	var err error

	if file, err = os.Open("config.toml"); err != nil {
		panic("Unable to load config.toml")
	}

	if data, err = ioutil.ReadAll(file); err != nil {
		panic("Unable to load config.toml")
	}

	var app Application
	if _, err = toml.Decode(string(data), &app); err != nil {
		panic("Unable to load config.toml")
	}

	return app
}
