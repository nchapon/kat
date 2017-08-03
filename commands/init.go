package commands

import (
	"fmt"
)

// setenv ...
func Init() {
	fmt.Println("Init Kat Project")

	config := Config{KongUrl: "https://localhost:8443",
		KongAdminUrl: "http://localhost:8001"}

	config.Save()

	fmt.Printf("✔ Initialized empty Kat config in %s \n", ConfigFile)
}
