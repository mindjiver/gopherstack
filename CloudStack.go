package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"flag"
)

type CloudStackClient struct {
	// The http client for communicating
	client *http.Client

	// The base URL of the API
	BaseURL string

	// Credentials
	APIKey string
	Secret string
}

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

func NewRequest(cloudstack CloudStackClient, request string) {
	var args = make(map[string]string)
	args["apikey"] = cloudstack.APIKey
	args["command"] = request
	args["response"] = "json"

	// we need to create the URL with a list of (key, value) of
	// arguments sorted in alphabetical order.
	var keys []string
	for k, _ := range args {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Create the initial api call string.
	var params []string
	for _, k := range keys {
		params = append(params, (k + "=" + url.QueryEscape(args[k])))
	}

	s := strings.Join(params, "&")

	// Generate signature for API call
	// * Convert the entire argument string to lowercase
	// * Calculate HMAC SHA1 of argument string with CloudStack secret
	// * URL encode the string and convert to base64
	var s2 = strings.ToLower(s)
	mac := hmac.New(sha1.New, []byte(cloudstack.Secret))
	mac.Write([]byte(s2))
	signature := base64.URLEncoding.EncodeToString(mac.Sum(nil))
	signature = url.QueryEscape(signature)
	// Apparently we need to manually(?) escape the underscore
	signature = strings.Replace(signature, "_", "%2F", -1)

	// Create the final URL before we issue the request
	api_call := cloudstack.BaseURL + "?" + s + "&signature=" + signature

	fmt.Println("Calling: " + api_call)

	// Print the results if we recieve a 200 response.
	resp, err := cloudstack.client.Get(api_call)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}

		fmt.Printf("%s\n", string(contents))
	}
}

func main() {
	var request = flag.String("command", "listVirtualMachines", "List Virtual Machines")
	flag.Parse()

	apiurl := os.Getenv("CLOUDSTACK_API_URL")
	if len(apiurl) == 0 {
		fmt.Println("Needed environment variable CLOUDSTACK_API_URL not found, exiting")
		os.Exit(1)
	}
	apikey := os.Getenv("CLOUDSTACK_API_KEY")
	if len(apikey) == 0 {
		fmt.Println("Needed environment variable CLOUDSTACK_API_KEY not found, exiting")
		os.Exit(1)
	}
	secret := os.Getenv("CLOUDSTACK_SECRET")
	if len(secret) == 0 {
		fmt.Println("Needed environment variable CLOUDSTACK_SECRET not found, exiting")
		os.Exit(1)
	}

	cs := CloudStackClient{}.New(apiurl, apikey, secret)
	NewRequest(*cs, *request)
}
