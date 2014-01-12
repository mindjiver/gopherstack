package gopherstack

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Creates a new client for communicating with CloudStack
func (cloudstack CloudStackClient) New(apiurl string, apikey string, secret string) *CloudStackClient {
	c := &CloudStackClient{
		client: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
			},
		},
		BaseURL: apiurl,
		APIKey:  apikey,
		Secret:  secret,
	}
	return c
}

func NewRequest(c CloudStackClient, request string, params url.Values) (interface{}, error) {
	client := c.client

	params.Set("apikey", c.APIKey)
	params.Set("command", request)
	params.Set("response", "json")

	// Generate signature for API call
	// * Serialize parameters and sort them by key, done by Encode
	// * Convert the entire argument string to lowercase
	// * Calculate HMAC SHA1 of argument string with CloudStack secret
	// * URL encode the string and convert to base64
	s := params.Encode()
	s2 := strings.ToLower(s)
	mac := hmac.New(sha1.New, []byte(c.Secret))
	mac.Write([]byte(s2))
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	signature = url.QueryEscape(signature)

	// Create the final URL before we issue the request
	url := c.BaseURL + "?" + s + "&signature=" + signature

	log.Printf("Calling %s ", url)

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	log.Printf("response from cloudstack: %d - %s", resp.StatusCode, body)
	if resp.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("Received HTTP client/server error from CloudStack: %d", resp.StatusCode))
		return nil, err
	}

	switch request {
	default:
		log.Printf("Unknown request %s", request)
	case "createSSHKeyPair":
		var decodedResponse CreateSshKeyPairResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	case "deleteSSHKeyPair":
		var decodedResponse DeleteSshKeyPairResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	case "deployVirtualMachine":
		var decodedResponse DeployVirtualMachineResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	case "destroyVirtualMachine":
		var decodedResponse DestroyVirtualMachineResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	case "stopVirtualMachine":
		var decodedResponse StopVirtualMachineResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	case "listVirtualMachines":
		var decodedResponse ListVirtualMachinesResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	case "listProjects":
		var decodedResponse ListProjectsResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	case "listVolumes":
		var decodedResponse ListVolumesResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	case "createTemplate":
		var decodedResponse CreateTemplateResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	case "queryAsyncJobResult":
		var decodedResponse QueryAsyncJobResultResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	}

	// only reached with unknown request
	return "", nil
}

type CloudStackClient struct {
	// The http client for communicating
	client *http.Client

	// The base URL of the API
	BaseURL string

	// Credentials
	APIKey string
	Secret string
}
