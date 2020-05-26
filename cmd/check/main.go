package main

import (
	"log"

	rlog "github.com/digitalocean/concourse-resource-library/log"
	rshared "github.com/digitalocean/concourse-resource-library/resource"
	resource "github.com/digitalocean/github-resource"
)

func main() {
	input := rlog.WriteStdin()
	defer rlog.Close()

	var request rshared.CheckRequest
	err := request.Read(input)
	if err != nil {
		log.Fatalf("failed to read request input: %s", err)
	}

	response, err := resource.Check(request)
	if err != nil {
		log.Fatalf("failed to perform check: %s", err)
	}

	err = response.Write()
	if err != nil {
		log.Fatalf("failed to write response to stdout: %s", err)
	}

	log.Println("check complete")
}
