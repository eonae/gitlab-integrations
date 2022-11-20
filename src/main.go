package main

import (
	"fmt"
	"gi/src/groups"
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

	gid, err := groups.GetGroupId(client, "v3")
	if err != nil {
		log.Fatal(err)
	}

	proj, err := projects.Fetch(client, projects.ProjectFilters{
		GroupId: gid,
		Pattern: "^(.*-service|web-arm|terminal-bff|medpoint24-bot)$",
	})

	if err != nil {
		log.Fatal(err)
	}

	for _, p := range proj {
		fmt.Printf("%d - %s\n", p.ID, p.Path)
	}

	err = projects.WriteJson(filepath.Join(config.OutDir, "projects.json"), proj)
	if err != nil {
		log.Fatal()
	}
}
