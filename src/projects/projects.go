// Manipulating gitlab projects
package projects

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"github.com/xanzy/go-gitlab"
)

type ProjectFilters struct {
	GroupId int

	// Project name regexp pattern
	Pattern string
}

// Fetch Gets all gitlab projects into RAM
func Fetch(client *gitlab.Client, filters ProjectFilters) ([]*gitlab.Project, error) {
	page := 1
	projects := make([]*gitlab.Project, 0)

	for {
		fmt.Printf("Getting projects page = %d\n", page)
		fetched, response, err := client.Groups.ListGroupProjects(filters.GroupId, &gitlab.ListGroupProjectsOptions{
			Simple: gitlab.Bool(true),
			ListOptions: gitlab.ListOptions{
				PerPage: 100,
				Page:    page,
			},
		})

		if err != nil {
			return nil, err
		}

		for _, project := range fetched {
			match := regexp.
				MustCompile(filters.Pattern).
				Match([]byte(project.Path))

			if match {
				projects = append(projects, project)
			} else {
				fmt.Printf("Filtered: %s\n", project.Path)
			}
		}

		fmt.Printf("Fetched projectes %d/%d\n", len(projects), response.TotalItems)
		page = response.NextPage

		if page == 0 {
			return projects, nil
		}
	}
}

// WriteJson writes a collection of projects to json file
func WriteJson(filename string, projects []*gitlab.Project) error {
	data, err := json.MarshalIndent(projects, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0744)
}
