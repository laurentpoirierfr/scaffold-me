package main

import (
	"flag"
	"log"
	"os"

	"github.com/laurentpoirierfr/scaffold-me/scaffold"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

const (
	DEFAULT_SCAFFOLD_URL     = "https://github.com/laurentpoirierfr/default-scaffold.git"
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
	checkIfError(err)
	defer os.Remove(dname)

	var referenceName plumbing.ReferenceName

	if *version == DEFAULT_SCAFFOLD_VERSION {
		referenceName = plumbing.NewBranchReferenceName(*version)
	} else {
		referenceName = plumbing.NewTagReferenceName(*version)
	}

	_, err = git.PlainClone(dname, false, &git.CloneOptions{
		URL:               *repo,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		ReferenceName:     referenceName,
		SingleBranch:      true,
	})

	checkIfError(err)
	os.RemoveAll(dname + "/.git")

	path, err := os.Getwd()
	checkIfError(err)

	sourceFolder := dname
	targetFolder := path

	scaffolder, err := scaffold.NewScaffolder(sourceFolder, targetFolder)
	checkIfError(err)
	err = scaffolder.Execute()
	checkIfError(err)
	log.Println("Bye ...")
}

func checkIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
