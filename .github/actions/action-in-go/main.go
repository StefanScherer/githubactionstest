package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sethvargo/go-githubactions"
)

func main() {
	fruit := githubactions.GetInput("fruit")
	if fruit == "" {
		githubactions.Fatalf("missing input 'fruit'")
	}
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")

	fmt.Println("Got access key: ", accessKey)

	jsonFile := githubactions.GetInput("jsonFile")
	if jsonFile == "" {
		githubactions.Fatalf("missing input 'jsonFile'")
	}

	_, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		exitErrorf("Error reading json file %s, %v", jsonFile, err)
	}

	fmt.Println("Done!")
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
