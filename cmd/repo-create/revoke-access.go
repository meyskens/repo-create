package main

import (
	"context"
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-github/v29/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

func init() {
	rootCmd.AddCommand(NewRevokeAccessCmd())
}

type revokeaccessCmdOptions struct {
	Branch       string
	MinReviewers int
}

// NewRevokeAccessCmd generates the `protect` command
func NewRevokeAccessCmd() *cobra.Command {
	s := revokeaccessCmdOptions{}
	c := &cobra.Command{
		Use:   "revoke-access",
		Short: "revoke all access of repositories",
		Long:  `Mass revoke all access of multiple repositories`,
		RunE:  s.RunE,
	}

	c.MarkFlagRequired("prefix")
	c.MarkFlagRequired("org")
	c.MarkFlagRequired("number")

	viper.BindPFlags(c.Flags())

	return c
}

func (s *revokeaccessCmdOptions) RunE(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: authToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	gh := github.NewClient(tc)

	for i := start; i <= number; i++ {
		name := fmt.Sprintf("%s%02d", prefix, i)
		collaborators, _, err := gh.Repositories.ListCollaborators(ctx, org, name, &github.ListCollaboratorsOptions{})
		spew.Dump(collaborators)
		if err != nil {
			return err
		}

		for _, collacollaborator := range collaborators {
			_, err := gh.Repositories.RemoveCollaborator(ctx, org, name, collacollaborator.GetLogin())
			if err != nil {
				log.Println(err)
			}
		}
	}

	return nil
}
