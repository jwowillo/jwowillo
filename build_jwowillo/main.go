package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jwowillo/jwowillo/project"
)

func main() {
	root, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	for repo := range project.Map {
		fmt.Println("Building repository at", repo)
		if err := project.BuildRepo(root, repo); err != nil {
			log.Fatal(err)
		}
	}
}
