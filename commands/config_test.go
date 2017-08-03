package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

var configFile string

// TestMain will exec each test, one by one
func TestMain(m *testing.M) {
	// exec setUp function
	setUp()
	// exec test and this returns an exit code to pass to os
	retCode := m.Run()
	// exec tearDown function
	tearDown()
	// If exit code is distinct of zero,
	// the test will be failed (red)
	os.Exit(retCode)
}

func setUp() {
	configFile = fmt.Sprintf("%s/.kat", katDir())
}

func tearDown() {
	err := os.Remove(configFile)
	if err != nil {
		panic(err)
	}

}

func TestSaveConfig(t *testing.T) {

	actual := Config{KongUrl: "https://localhost:8443",
		KongAdminUrl: "http://localhost:8001"}

	actual.Save()

	raw, err := ioutil.ReadFile(configFile)
	if err != nil {
		t.Errorf("Could not read file  %s", configFile)
		t.FailNow()
	}

	var expected Config
	json.Unmarshal(raw, &expected)

	if actual != expected {
		t.Errorf("Actual is not equal to expected !!")
	}

}

func TestGetConfig(t *testing.T) {

	savedConfig := Config{KongUrl: "https://localhost:8443",
		KongAdminUrl: "http://localhost:8001"}
	savedConfig.Save()

	readedConfig, _ := Get()

	if savedConfig != readedConfig {
		t.Errorf("Saved config and readed config are not equal !!")
	}

}
