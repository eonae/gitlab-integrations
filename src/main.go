package main

import (
	"fmt"
	"gi/src/projects"
	"log"

	"github.com/xanzy/go-gitlab"
)

func main() {
	config := mustParse()

	client, err := gitlab.NewClient(config.Token, gitlab.WithBaseURL(config.Url))
	if err != nil {
		log.Fatal(err)
	}

	for proj := range projects.Iter(client, 10) {
		fmt.Printf("project: %s\n", proj.Name)
	}

	// err = projects.WriteJson(filepath.Join(config.OutDir, "projects.json"), proj)
	// if err != nil {
	// 	log.Fatal()
	// }
}
