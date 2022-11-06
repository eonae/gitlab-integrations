// Manipulating gitlab projects
package projects

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/xanzy/go-gitlab"
)

// Iter Gets all gitlab projects into RAM
func Iter(client *gitlab.Client, buffer int) chan *gitlab.Project {
	output := make(chan *gitlab.Project, buffer)

	go func() {
		page := 1
		projects := make([]*gitlab.Project, 0)

		for {
			fmt.Printf("Getting projects page = %d\n", page)
			fetched, response, err := client.Projects.ListProjects(&gitlab.ListProjectsOptions{
				// WTF?? Why pointer to bool? Ah maybe to accept nil!
				Simple: gitlab.Bool(false),
				ListOptions: gitlab.ListOptions{
					PerPage: buffer / 2,
					Page:    page,
				},
			})

			if err != nil {
				close(output)
				break
			}

			projects = append(projects, fetched...)

			fmt.Printf("Fetched projectes %d/%d\n", len(projects), response.TotalItems)

			for _, p := range projects {
				output <- p
			}

			page = response.NextPage

			if page == 0 {
				close(output)
				break
			}
		}
	}()

	return output
}

// WriteJson writes a collection of projects to json file
func WriteJson(filename string, projects []*gitlab.Project) error {
	data, err := json.MarshalIndent(projects, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0744)
}
