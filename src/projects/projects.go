// Manipulating gitlab projects
package projects

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/xanzy/go-gitlab"
)

// FetchAll Gets all gitlab projects into RAM
func FetchAll(client *gitlab.Client) ([]*gitlab.Project, error) {
	page := 1
	projects := make([]*gitlab.Project, 0)

	for {
		fmt.Printf("Getting projects page = %d\n", page)
		fetched, response, err := client.Projects.ListProjects(&gitlab.ListProjectsOptions{
			// WTF?? Why pointer to bool? Ah maybe to accept nil!
			Simple: gitlab.Bool(false),
			ListOptions: gitlab.ListOptions{
				PerPage: 100,
				Page:    page,
			},
		})

		if err != nil {
			return nil, err
		}

		projects = append(projects, fetched...)

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
