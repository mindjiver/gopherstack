package gopherstack

import (
	"net/url"
)

func (c CloudStackClient) ListDiskOfferings(domainid string, id string, keyword string, name string, page string, pagesize string) (string, error) {
	params := url.Values{}
	//params.Set("domainid", domainid)
	_, err := NewRequest(c, "listDiskOfferings", params)
	if err != nil {
		return "", err
	}
	return "", err
}

type DiskOffering struct {
	Created      string  `json:"created"`
	Disksize     float64 `json:"disksize"`
	Displaytext  string  `json:"displaytext"`
	ID           string  `json:"id"`
	Iscustomized bool    `json:"iscustomized"`
	Name         string  `json:"name"`
	Storagetype  string  `json:"storagetype"`
}

type ListDiskOfferingsResponse struct {
	Listdiskofferingsresponse struct {
		Count        float64        `json:"count"`
		Diskoffering []DiskOffering `json:"diskoffering"`
	} `json:"listdiskofferingsresponse"`
}
