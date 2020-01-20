package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v29/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

func init() {
	rootCmd.AddCommand(NewProjectCmd())
}

type projectCmdOptions struct {
	Prefix   string
	Number   int
	Org      string
	Name     string
	Collumns []string
}

// NewProjectCmd generates the `project` command
func NewProjectCmd() *cobra.Command {
	s := projectCmdOptions{}
	c := &cobra.Command{
		Use:   "project",
		Short: "add project to repositories",
		Long:  `Mass add a project to multiple repositories`,
		RunE:  s.RunE,
	}
	c.Flags().StringVarP(&s.Prefix, "prefix", "p", "", "Prefix of repository names")
	c.Flags().IntVarP(&s.Number, "number", "n", 1, "How many repositories to label")
	c.Flags().StringVarP(&s.Org, "org", "o", "", "Org to label repos under")

	c.Flags().StringVar(&s.Name, "name", "", "Name for the project")
	c.Flags().StringArrayVar(&s.Collumns, "collumns", []string{}, "Collumns for the project")

	c.MarkFlagRequired("prefix")
	c.MarkFlagRequired("org")
	c.MarkFlagRequired("number")

	c.MarkFlagRequired("name")

	viper.BindPFlags(c.Flags())

	return c
}

func (s *projectCmdOptions) RunE(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: authToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	gh := github.NewClient(tc)

	for i := 1; i <= s.Number; i++ {
		name := fmt.Sprintf("%s-A%02d", s.Prefix, i)

		proj, _, err := gh.Repositories.CreateProject(ctx, s.Org, name, &github.ProjectOptions{
			Name: &s.Name,
		})
		if err != nil {
			log.Println(err)
		}

		for _, collumn := range s.Collumns {
			_, _, err := gh.Projects.CreateProjectColumn(ctx, proj.GetID(), &github.ProjectColumnOptions{Name: collumn})
			if err != nil {
				log.Println(err)
			}
		}

	}

	return nil
}
