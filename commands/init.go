package commands

import (
	"encoding/json"
	"fmt"
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

// setenv ...
func Init() {
	fmt.Println("Init Kat Project")

	config := Config{KongUrl: "https://localhost:8443",
		KongAdminUrl: "http://localhost:8001"}

	json, _ := json.MarshalIndent(config, "", "  ")

	ioutil.WriteFile(ConfigFile, json, os.ModePerm)

	fmt.Printf("✔ Initialized empty Kat config in %s \n", ConfigFile)
}
