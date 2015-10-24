package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Config struct {
	StaticPaths    []string `json:"static_paths"`
	TemplatePaths  []string `json:"templates"`
	SiteIndexTitle string   `json:"site_index_title"`

	PathConfig string `json:"-"`
}

func FromConfiguration(path string) Config {
	conf := Config{}
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(b, &conf); err != nil {
		panic(err)
	}
	conf.PathConfig = path

	sanitizePaths(conf)
	return conf
}

func sanitizePaths(conf Config) {
	for i, path := range conf.StaticPaths {
		// try the path as given first
		if _, err := os.Stat(path); err != nil {
			// now check if they are relative to the config file
			confBase := filepath.Dir(conf.PathConfig)
			if _, err := os.Stat(filepath.Join(confBase, path)); err != nil {
				panic(err)
			}
			conf.StaticPaths[i] = filepath.Join(confBase, path)
		}
	}
}
