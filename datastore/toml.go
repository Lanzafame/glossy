package datastore

import (
	"github.com/lanzafame/glossy/conf"
)

// tomlStorePath represents the environment variable
var tomlStorePath = "GLSY_TOML_PATH"

type TomlStore struct {
	FilePath string
}

func (ts *TomlStore) Write(bs []byte) error {
	return nil
}

func (ts *TomlStore) Read() ([]byte, error) {
	return nil, nil
}

func NewTomlStore(conf conf.Config) (Store, error) {
	fp := conf.Get(tomlStorePath, "~/.glsy.toml")

	return &TomlStore{
		FilePath: fp,
	}, nil
}
