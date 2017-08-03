package commands

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

var ConfigFile string

// init ...
func init() {

	katDir := katDir()
	ConfigFile = filepath.Join(katDir, ".kat")

	err := os.MkdirAll(katDir, os.ModePerm)

	check(err)

	if _, err := os.Stat(ConfigFile); os.IsNotExist(err) {
		os.Create(ConfigFile)
	}

}

// katDir ...
func katDir() string {
	katDir, err := os.Getwd()

	check(err)

	return katDir
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Config struct {
	KongUrl      string
	KongAdminUrl string
}

// Save ...
func (c Config) Save() error {

	json, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(ConfigFile, json, os.ModePerm)
}

func Get() (Config, error) {

	raw, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(err)
	}

	c := &Config{}

	err = json.Unmarshal(raw, &c)

	if err != nil {
		return *c, err
	}

	return *c, nil

}
