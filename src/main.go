package main

import (
	"gi/src/projects"
	"log"
	"path/filepath"

	"github.com/xanzy/go-gitlab"
)

func main() {
	config := mustParse()

	client, err := gitlab.NewClient(config.Token, gitlab.WithBaseURL(config.Url))
	if err != nil {
		log.Fatal(err)
	}

	proj, err := projects.FetchAll(client)
	if err != nil {
		log.Fatal(err)
	}

	err = projects.WriteJson(filepath.Join(config.OutDir, "projects.json"), proj)
	if err != nil {
		log.Fatal()
	}
}
