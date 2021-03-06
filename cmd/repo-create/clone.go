package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(NewCloneCmd())
}

type cloneCmdOptions struct {
	Source string
	Force  bool
}

// NewCloneCmd generates the `clone` command
func NewCloneCmd() *cobra.Command {
	s := cloneCmdOptions{}
	c := &cobra.Command{
		Use:   "clone",
		Short: "Clones a repo into multiple repositories",
		Long:  `Mass clone a repo and put it in newly created repositories`,
		RunE:  s.RunE,
	}
	c.Flags().StringVarP(&s.Source, "source", "f", "", "Base repository")
	c.Flags().BoolVar(&s.Force, "force", false, "Force push !!DANGEROUS!!")

	c.MarkFlagRequired("prefix")
	c.MarkFlagRequired("org")
	c.MarkFlagRequired("number")
	c.MarkFlagRequired("source")

	viper.BindPFlags(c.Flags())

	return c
}

func (s *cloneCmdOptions) RunE(cmd *cobra.Command, args []string) error {
	dir, err := ioutil.TempDir(os.TempDir(), "repo-create-clone")
	if err != nil {
		return err
	}

	log.Println("Cloning to", dir)

	r, err := git.PlainClone(dir, false, &git.CloneOptions{
		// The intended use of a GitHub personal access token is in replace of your password
		// because access tokens can easily be revoked.
		// https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/
		Auth: &http.BasicAuth{
			Username: "iloveoctocats", // yes, this can be anything except an empty string
			Password: authToken,
		},
		URL: s.Source,
	})

	if err != nil {
		return err
	}

	for i := start; i <= number; i++ {
		name := fmt.Sprintf("%s%02d", prefix, i)
		log.Println("pushing", name)

		_, err := r.CreateRemote(&config.RemoteConfig{
			Name: name,
			URLs: []string{fmt.Sprintf("https://github.com/%s/%s.git", org, name)},
		})

		if err != nil {
			log.Println(err)
		}

		pushOptions := git.PushOptions{
			RemoteName: name,
			// The intended use of a GitHub personal access token is in replace of your password
			// because access tokens can easily be revoked.
			// https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/
			Auth: &http.BasicAuth{
				Username: "iloveoctocats", // yes, this can be anything except an empty string
				Password: authToken,
			},
		}
		if s.Force {
			pushOptions.RefSpecs = []config.RefSpec{"+refs/heads/master:refs/heads/master"}
		}
		err = r.Push(&pushOptions)

		if err != nil {
			log.Println(err)
		}

	}

	return nil
}
