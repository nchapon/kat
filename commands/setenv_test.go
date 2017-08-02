package commands

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	t.Log("Testing config file is created.")

	Init()

	configFile := fmt.Sprintf("%s/kong.conf", katDir())

	f, err := os.Open(configFile)

	if err != nil {
		t.Errorf("Could not read file  %s", configFile)
	}

	// TODO : change tesy
	s := bufio.NewScanner(f)
	s.Scan()
	if data := s.Text(); data != "test" {
		t.Errorf("Expected content=test, but it was %s.", data)
	}

}
