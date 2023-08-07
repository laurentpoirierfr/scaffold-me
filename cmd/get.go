/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"github.com/laurentpoirierfr/scaffold-me/scaffold"
	"github.com/laurentpoirierfr/scaffold-me/util"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get and execute scaffolder from url.",
	Long:  "Get and execute scaffolder from url.",
	Run: func(cmd *cobra.Command, args []string) {

		branch, err := cmd.Flags().GetString("branch")
		util.CheckIfError(err)
		tag, err := cmd.Flags().GetString("tag")
		util.CheckIfError(err)
		url, err := cmd.Flags().GetString("url")
		util.CheckIfError(err)

		user, err := cmd.Flags().GetString("user")
		util.CheckIfError(err)
		password, err := cmd.Flags().GetString("password")
		util.CheckIfError(err)

		dname, err := os.MkdirTemp("", "scaffold")
		util.CheckIfError(err)
		defer os.Remove(dname)

		var referenceName plumbing.ReferenceName

		if tag == "" {
			referenceName = plumbing.NewBranchReferenceName(branch)
		} else {
			referenceName = plumbing.NewTagReferenceName(tag)
		}

		if user != "" {
			_, err = git.PlainClone(dname, false, &git.CloneOptions{
				URL:               url,
				RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
				ReferenceName:     referenceName,
				SingleBranch:      true,
				Auth: &http.BasicAuth{
					Username: user, // yes, this can be anything except an empty string
					Password: password,
				},
			})
		} else {
			_, err = git.PlainClone(dname, false, &git.CloneOptions{
				URL:               url,
				RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
				ReferenceName:     referenceName,
				SingleBranch:      true,
			})
		}
		util.CheckIfError(err)
		os.RemoveAll(dname + "/.git")

		path, err := os.Getwd()
		util.CheckIfError(err)

		sourceFolder := dname
		targetFolder := path

		scaffolder, err := scaffold.NewScaffolder(sourceFolder, targetFolder)
		util.CheckIfError(err)
		err = scaffolder.Execute()
		util.CheckIfError(err)
		log.Println("Bye ...")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	getCmd.PersistentFlags().String("url", "https://github.com/laurentpoirierfr/default-scaffold.git", "Git url off scaffolder.")
	getCmd.PersistentFlags().String("branch", "main", "Branch name off scaffolder.")
	getCmd.PersistentFlags().String("tag", "", "Tag version off scaffolder.")

	getCmd.PersistentFlags().String("user", "", "User login to git repository.")
	getCmd.PersistentFlags().String("password", "", "User password to git repository.")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
