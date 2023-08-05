package main

import (
	"flag"
	"log"
	"os"

	"github.com/laurentpoirierfr/scaffold-me/scaffold"
	"gopkg.in/src-d/go-git.v4"
)

const (
	DEFAULT_SCAFFOLD_URL     = "https://github.com/laurentpoirierfr/default-scaffold"
	DEFAULT_SCAFFOLD_VERSION = "main"
)

func main() {

	repo := flag.String("r", "", "Git repository where scaffold is located.")
	version := flag.String("v", "main", "Version of git repository.")
	flag.Parse()

	if *repo == "" {
		*repo = DEFAULT_SCAFFOLD_URL
	}

	if *version == "" {
		*version = DEFAULT_SCAFFOLD_VERSION
	}

	dname, err := os.MkdirTemp("", "scaffold")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(dname)

	_, err = git.PlainClone(dname, false, &git.CloneOptions{
		URL:               *repo,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})
	if err != nil {
		log.Fatal(err)
	}
	os.RemoveAll(dname + "/.git")

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	sourceFolder := dname
	targetFolder := path

	scaffolder, err := scaffold.NewScaffolder(sourceFolder, targetFolder)
	if err != nil {
		log.Fatal(err)
	}
	err = scaffolder.Execute()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Bye ...")
	}
}
