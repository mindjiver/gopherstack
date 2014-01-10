package cloudstack

import (
	"net/url"
)

type Template struct {
	Id   string
	Name string
}

type TemplatesResponse struct {
	Templates []Template
}

// Creates a Template of a Virtual Machine by it's ID
func (c CloudStackClient) CreateTemplate(displaytext string, name string, volumeid string, ostypeid string) (string, error) {
	params := url.Values{}
	params.Set("displaytext", displaytext)
	params.Set("name", name)
	params.Set("ostypeid", ostypeid)
	params.Set("volumeid", volumeid)
	_, err := NewRequest(c, "createTemplate", params)
	// return async job id
	return "jobId", err
}

// Returns all available templates
func (c CloudStackClient) Templates() ([]Template, error) {
	params := url.Values{}
	_, err := NewRequest(c, "listTemplates", params)
	// unmarshall json to a proper list
	return nil, err
}

// Deletes an template by its ID.
func (c CloudStackClient) DeleteTemplate(id string) (uint, error) {
	params := url.Values{}
	params.Set("id", id)
	_, err := NewRequest(c, "deleteTemplate", params)
	return 0, err
}

