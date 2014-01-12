package gopherstack

import (
	"net/url"
)

// List volumes for Virtual Machine by it's ID
func (c CloudStackClient) ListVolumes(vmid string) (string, error) {
	params := url.Values{}
	params.Set("virtualmachineid", vmid)
	response, err := NewRequest(c, "listVolumes", params)
	if err != nil {
		return "", err
	}

	count := response.(ListVolumesResponse).Listvolumesresponse.Count
	// if there are no volumes we just return here
	if count < 1 {
		return "", err
	}

	volumeId := response.(ListVolumesResponse).Listvolumesresponse.Volume[0].ID

	return volumeId, err
}

type ListVolumesResponse struct {
	Listvolumesresponse struct {
		Count  float64 `json:"count"`
		Volume []struct {
			Account                    string        `json:"account"`
			Created                    string        `json:"created"`
			Destroyed                  bool          `json:"destroyed"`
			Deviceid                   float64       `json:"deviceid"`
			Domain                     string        `json:"domain"`
			Domainid                   string        `json:"domainid"`
			ID                         string        `json:"id"`
			Isextractable              bool          `json:"isextractable"`
			Name                       string        `json:"name"`
			Serviceofferingdisplaytext string        `json:"serviceofferingdisplaytext"`
			Serviceofferingid          string        `json:"serviceofferingid"`
			Serviceofferingname        string        `json:"serviceofferingname"`
			Size                       float64       `json:"size"`
			State                      string        `json:"state"`
			Storage                    string        `json:"storage"`
			Storagetype                string        `json:"storagetype"`
			Tags                       []interface{} `json:"tags"`
			Type                       string        `json:"type"`
			Virtualmachineid           string        `json:"virtualmachineid"`
			Vmdisplayname              string        `json:"vmdisplayname"`
			Vmname                     string        `json:"vmname"`
			Vmstate                    string        `json:"vmstate"`
			Zoneid                     string        `json:"zoneid"`
			Zonename                   string        `json:"zonename"`
		} `json:"volume"`
	} `json:"listvolumesresponse"`
}
