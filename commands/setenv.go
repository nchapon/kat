package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

type Env struct {
	name string
	dir  string
}

// NewGame ...
func NewEnv(n string) Env {
	return Env{name: n, dir: "/tmp"}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// setenv ...
func SetEnv(name string) {
	fmt.Println("Set env", name)

	ex, err := os.Executable()
	check(err)

	exPath := filepath.Dir(ex)

	f, err := os.Create(fmt.Sprintf("%s/.kat", exPath))
	check(err)
	defer f.Close()

	f.WriteString(fmt.Sprintln(name))
	f.Sync()

	fmt.Printf("Created .kat file in %s.\n", exPath)
}
