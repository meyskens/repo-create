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

	c.Flags().StringVar(&s.Name, "name", "", "Name for the project")
	c.Flags().StringArrayVar(&s.Collumns, "collumns", []string{}, "Collumns for the project")

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

	for i := 1; i <= number; i++ {
		name := fmt.Sprintf("%s%02d", prefix, i)

		proj, _, err := gh.Repositories.CreateProject(ctx, org, name, &github.ProjectOptions{
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
