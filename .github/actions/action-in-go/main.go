package main

import "github.com/sethvargo/go-githubactions"

func main() {
	fruit := githubactions.GetInput("fruit")
	if fruit == "" {
		githubactions.Fatalf("missing input 'fruit'")
	}
	accessKey := githubactions.GetInput("AWS_ACCESS_KEY_ID")
	if accessKey == "" {
		githubactions.Fatalf("missing input 'AWS_ACCESS_KEY_ID'")
	}
	githubactions.AddMask(accessKey)
	secretKey := githubactions.GetInput("AWS_SECRET_ACCESS_KEY")
	if secretKey == "" {
		githubactions.Fatalf("missing input 'AWS_SECRET_ACCESS_KEY'")
	}
	githubactions.AddMask(secretKey)
}
