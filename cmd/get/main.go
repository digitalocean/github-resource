package main

import (
	"log"

	rlog "github.com/digitalocean/concourse-resource-library/log"
	resource "github.com/digitalocean/github-resource"
)

func main() {
	input := rlog.WriteStdin()
	defer rlog.Close()

	var request resource.GetRequest
	err := request.Read(input)
	if err != nil {
		log.Fatalf("failed to read request input: %s", err)
	}

	response, err := resource.Get(request)
	if err != nil {
		log.Fatalf("failed to perform Get: %s", err)
	}

	err = response.Write()
	if err != nil {
		log.Fatalf("failed to write response to stdout: %s", err)
	}

	log.Println("Get complete")
}
