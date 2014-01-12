package gopherstack

import (
	"net/url"
)

// Deploys a Virtual Machine and returns it's id
func (c CloudStackClient) AttachIso(isoid string, vmid string) (string, error) {
	params := url.Values{}
	params.Set("id", isoid)
	params.Set("virtualmachineid", vmid)

	_, err := NewRequest(c, "attachIso", params)
	if err != nil {
		return "", err
	}
	//jobid := response.(AttachIsoResponse).Attachisoresponse.Jobid
	return "", err
}

func (c CloudStackClient) DetachIso(vmid string) (string, error) {
	params := url.Values{}
	params.Set("virtualmachineid", vmid)
	_, err := NewRequest(c, "detachIso", params)
	if err != nil {
		return "", err
	}
	//jobid := response.(DetachIsoResponse).DetachIsoresponse.Jobid
	return "", err
}

func (c CloudStackClient) ListIsos() (string, error) {
	_, err := NewRequest(c, "listIsos", nil)
	if err != nil {
		return "", err
	}
	//jobid := response.(ListIsosResponse).Listisosresponse.Jobid
	return "", err
}
