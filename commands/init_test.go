package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestInit(t *testing.T) {
	t.Log("Testing config file is created.")

	Init()

	configFile := fmt.Sprintf("%s/.kat", katDir())
	raw, err := ioutil.ReadFile(configFile)
	if err != nil {
		t.Errorf("Could not read file  %s", configFile)
		t.FailNow()
	}

	actual := Config{KongUrl: "https://localhost:8443",
		KongAdminUrl: "http://localhost:8001"}

	var expected Config
	json.Unmarshal(raw, &expected)

	if actual != expected {
		t.Errorf("Actual is not equal to expeted !!")
	}

}
