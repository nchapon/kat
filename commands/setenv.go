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
	ConfigFile = filepath.Join(katDir, "kong.conf")

	err := os.MkdirAll(katDir, os.ModePerm)

	check(err)

	if _, err := os.Stat(ConfigFile); os.IsNotExist(err) {
		os.Create(ConfigFile)
	}

}

// katDir ...
func katDir() string {
	exPath, err := os.Executable()
	check(err)

	return filepath.Dir(exPath)
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

	fmt.Printf("Your Kong is now %s \n", ConfigFile)
}
