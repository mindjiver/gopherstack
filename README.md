gopherstack
===========

Cloudstack API library written in Go. Only tested towards Cloudstack
3.0.6 so far. Main use so far has been to serve as a library for a
[Packer.io](http://www.packer.io) builder.

Example usage
-------------

Showing the IP and state of a virtual machine:

```go
package main

import (
	"fmt"
	"github.com/mindjiver/gopherstack"
	"os"
)

func main() {

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

	cs := gopherstack.CloudStackClient{}.New(apiurl, apikey, secret)

	vmid := "19d2acfb-e281-4a13-8d62-e04ab501271d"
	response, err := cs.ListVirtualMachines(vmid)
	if err != nil {
		fmt.Errorf("Error listing virtual machine: %s", err)
		os.Exit(1)
	}
	
	if len(response.Listvirtualmachinesresponse.Virtualmachine) > 0 {
		ip := response.Listvirtualmachinesresponse.Virtualmachine[0].Nic[0].Ipaddress
		state := response.Listvirtualmachinesresponse.Virtualmachine[0].State
		fmt.Printf("%s has IP : %s and state : %s\n", vmid, ip, state)
	} else {
		fmt.Printf("No VM with UUID: %s found\n", vmid)
	}

}
```

[![Travis CI status](https://travis-ci.org/mindjiver/gopherstack.png?branch=master)](https://travis-ci.org/mindjiver/gopherstack/builds/)
[![GoDoc](https://godoc.org/github.com/mindjiver/gopherstack?status.png)](https://godoc.org/github.com/mindjiver/gopherstack)
