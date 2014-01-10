package gopherstack

import (
	"net/url"
)

// List the available CloudStack projects
func (c CloudStackClient) ListProjects(name string) (string, error) {
	params := url.Values{}

	if name != "" {
		params.Set("name", name)
	}
	_, err := NewRequest(c, "listProjects", params)
	if err != nil {
		return "", err
	}

	return "", err
}
