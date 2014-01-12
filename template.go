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
func (c CloudStackClient) ListTemplates(name string, filter string) (string, string, error) {
	params := url.Values{}
	params.Set("name", name)
	params.Set("templatefilter", filter)
	response, err := NewRequest(c, "listTemplates", params)
	if err != nil {
		return "", "", err
	}

	count := response.(ListTemplatesResponse).Listtemplatesresponse.Count
	if count < 1 {
		return "", "", err
	}

	foundName := response.(ListTemplatesResponse).Listtemplatesresponse.Template[0].Name
	templateId := response.(ListTemplatesResponse).Listtemplatesresponse.Template[0].ID
	return foundName, templateId, err
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

type Template struct {
	Account          string        `json:"account"`
	Created          string        `json:"created"`
	CrossZones       bool          `json:"crossZones"`
	Displaytext      string        `json:"displaytext"`
	Domain           string        `json:"domain"`
	Domainid         string        `json:"domainid"`
	Format           string        `json:"format"`
	Hypervisor       string        `json:"hypervisor"`
	ID               string        `json:"id"`
	Isextractable    bool          `json:"isextractable"`
	Isfeatured       bool          `json:"isfeatured"`
	Ispublic         bool          `json:"ispublic"`
	Isready          bool          `json:"isready"`
	Name             string        `json:"name"`
	Ostypeid         string        `json:"ostypeid"`
	Ostypename       string        `json:"ostypename"`
	Passwordenabled  bool          `json:"passwordenabled"`
	Size             float64       `json:"size"`
	Sourcetemplateid string        `json:"sourcetemplateid"`
	Sshkeyenabled    bool          `json:"sshkeyenabled"`
	Status           string        `json:"status"`
	Tags             []interface{} `json:"tags"`
	Templatetype     string        `json:"templatetype"`
	Zoneid           string        `json:"zoneid"`
	Zonename         string        `json:"zonename"`
}

type ListTemplatesResponse struct {
	Listtemplatesresponse struct {
		Count    float64    `json:"count"`
		Template []Template `json:"template"`
	} `json:"listtemplatesresponse"`
}
