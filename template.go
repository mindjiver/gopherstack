package gopherstack

import (
	"net/url"
)

// Creates a Template of a Virtual Machine by it's ID
func (c CloudStackClient) CreateTemplate(displaytext string, name string, volumeid string, ostypeid string) (string, error) {
	params := url.Values{}
	params.Set("displaytext", displaytext)
	params.Set("name", name)
	params.Set("ostypeid", ostypeid)
	params.Set("volumeid", volumeid)

	response, err := NewRequest(c, "createTemplate", params)
	if err != nil {
		return "", err
	}

	jobId := response.(CreateTemplateResponse).Createtemplateresponse.Jobid
	return jobId, err
}

// Returns all available templates
func (c CloudStackClient) ListTemplates(name string) ([]string, error) {
	params := url.Values{}
	params.Set("name", name)
	_, err := NewRequest(c, "listTemplates", params)
	if err != nil {
		return nil, err
	}

	return nil, err
}

// Deletes an template by its ID.
func (c CloudStackClient) DeleteTemplate(id string) (string, error) {
	params := url.Values{}
	params.Set("id", id)
	_, err := NewRequest(c, "deleteTemplate", params)
	if err != nil {
		return "", err
	}
	return "", err
}

type CreateTemplateResponse struct {
	Createtemplateresponse struct {
		ID    string `json:"id"`
		Jobid string `json:"jobid"`
	} `json:"createtemplateresponse"`
}
