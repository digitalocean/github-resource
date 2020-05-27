package main

import (
	"log"

	rlog "github.com/digitalocean/concourse-resource-library/log"
	resource "github.com/digitalocean/github-resource"
)

func main() {
	input := rlog.WriteStdin()
	defer rlog.Close()

	var request resource.PutRequest
	err := request.Read(input)
	if err != nil {
		log.Fatalf("failed to read request input: %s", err)
	}

	response, err := resource.Put(request)
	if err != nil {
		log.Fatalf("failed to perform Put: %s", err)
	}

	err = response.Write()
	if err != nil {
		log.Fatalf("failed to write response to stdout: %s", err)
	}

	log.Println("Put complete")
}
