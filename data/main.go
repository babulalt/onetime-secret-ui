package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
)

type Dockerfile struct {
	Content string `json:"content"`
}

func main() {
	// Specify the path to the Dockerfile
	dockerfilePath := "/media/sugam/8E5AAD495AAD2EC1/workspace/src/github.com/babu-github/one-time-secret/Dockerfile"

	// Read the contents of the Dockerfile
	dockerfileContents, err := readDockerfile(dockerfilePath)
	if err != nil {
		fmt.Printf("Error reading Dockerfile: %s\n", err)
		return
	}

	// Display the Dockerfile contents as a string
	fmt.Println(dockerfileContents)
}

func readDockerfile(filePath string) (string, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Read the file contents
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	a := Dockerfile{}
	mp := map[string]interface{}{
		"content": string(contents),
	}
	josnBytes, _ := json.Marshal(mp)
	logrus.Info(string(josnBytes))
	_ = json.Unmarshal(josnBytes, &a)
	return string(a.Content), nil
}
