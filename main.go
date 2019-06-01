package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/machinebox/graphql"
	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	homeDir, err := homedir.Dir()
	mustNot(err)

	storyCfgPath := filepath.Join(homeDir, ".storyscript", "config")
	cfg := loadConfig(storyCfgPath)

	graph := graphql.NewClient("https://api.storyscript.io/graphql")
	client := client{
		graph: graph,
		token: cfg.AccessToken,
	}

	fmt.Println("collecting apps...")
	apps := client.AllApps()

	if len(os.Args) > 1 && os.Args[1] == "destroy" {
		client.DestroyApps(apps)
		os.Exit(0)
	}

	fmt.Println("Cowardly refusing to continue...")
	fmt.Println("Pass `destroy` as the first argument if you really want to destroy the following apps:")
	for _, app := range apps {
		fmt.Printf(" - %s\n", app.Name)
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

var mustNot = must
