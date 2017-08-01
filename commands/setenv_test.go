package commands

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestSetEnv(t *testing.T) {
	t.Logf("Testing .kat file is created in /tmp")

	//env := NewEnv("Test")

	SetEnv("test")

	ex, err := os.Executable()
	check(err)

	exPath := filepath.Dir(ex)

	configFile := fmt.Sprintf("%s/.kat", exPath)

	f, err := os.Open(configFile)

	if err != nil {
		t.Errorf("Could not read file  %s", configFile)
	}

	s := bufio.NewScanner(f)
	s.Scan()
	if data := s.Text(); data != "test" {
		t.Errorf("Expected content=test, but it was %s.", data)
	}
	//env.inc(5)

	//if current := env.current; current != 5 {
	//t.Errorf("Expected current=5, but it was %d.", current)
	//}

}
