package main

import (
	"fmt"

	"github.com/kminehart/caye"
)

func buildContainer(client *caye.Client, dockerfile string, image string, tag string) (*caye.Result, error) {
	return client.Run("docker",
		"build",
		".",
		"-f",
		dockerfile,
		"-t",
		fmt.Sprintf("%s:%s", image, tag),
	)
}

func main() {
	client := caye.Default()
	defer client.Done()

	version, err := client.Git().GetCommitSHA()
	if err != nil {
		client.Fatalln(err)
	}

	_, err = buildContainer(client, "Dockerfile", "376424775662.dkr.ecr.us-east-2.amazonaws.com/hyphen/api", version)
	if err != nil {
		client.Fatalln(err)
	}
}
