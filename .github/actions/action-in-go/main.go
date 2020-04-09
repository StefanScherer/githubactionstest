package main

import (
	"fmt"
	"os"

	"github.com/sethvargo/go-githubactions"
)

func main() {
	fruit := githubactions.GetInput("fruit")
	if fruit == "" {
		githubactions.Fatalf("missing input 'fruit'")
	}
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	fmt.Println("Got access key: ", accessKey)
	githubactions.AddMask(secretKey)
}
