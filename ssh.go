package cloudstack

import (
	"net/url"
)

// Create a SSH key pair
func (c CloudStackClient) CreateSSHKeyPair(name string) (string, error) {
	params := url.Values{}
	params.Set("name", name)
	response, err := NewRequest(c, "createSSHKeyPair", params)
	if err != nil {
		return "", err
	}
	privatekey := response.(CreateSshKeyPairResponse).Createsshkeypairresponse.Keypair.Privatekey
	return privatekey, nil
}

// Deletes an SSH key pair
func (c CloudStackClient) DeleteSSHKeyPair(name string) (string, error) {
	params := url.Values{}
	params.Set("name", name)
	response, err := NewRequest(c, "deleteSSHKeyPair", params)
	if err != nil {
		return "", err
	}
	success := response.(DeleteSshKeyPairResponse).Deletesshkeypairresponse.Success
	return success, err
}

type CreateSshKeyPairResponse struct {
	Createsshkeypairresponse struct {
		Keypair struct {
			Fingerprint string `json:"fingerprint"`
			Name        string `json:"name"`
			Privatekey  string `json:"privatekey"`
		} `json:"keypair"`
	} `json:"createsshkeypairresponse"`
}

type DeleteSshKeyPairResponse struct {
	Deletesshkeypairresponse struct {
		Success string `json:"success"`
	} `json:"deletesshkeypairresponse"`
}
