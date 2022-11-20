package groups

import (
	"fmt"

	"github.com/xanzy/go-gitlab"
)

func GetGroupId(client *gitlab.Client, groupName string) (int, error) {
	groups, _, err := client.Groups.SearchGroup("v3")

	if err != nil {
		return 0, err
	}

	if len(groups) == 0 {
		return 0, fmt.Errorf("No group matching %s found!", groupName)
	}

	if len(groups) > 1 {
		return 0, fmt.Errorf("More than one group matching %s found!", groupName)
	}

	return groups[0].ID, nil
}
