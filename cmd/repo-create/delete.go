package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v29/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

func init() {
	rootCmd.AddCommand(NewDeleteCmd())
}

type deleteCmdOptions struct {
	IsPrivate bool
}

// NewDeleteCmd generates the `delete` command
func NewDeleteCmd() *cobra.Command {
	s := deleteCmdOptions{}
	c := &cobra.Command{
		Use:   "delete",
		Short: "Deletes repositories",
		Long:  `Mass deletes repositories`,
		RunE:  s.RunE,
	}

	viper.BindPFlags(c.Flags())

	return c
}

func (s *deleteCmdOptions) RunE(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: authToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	gh := github.NewClient(tc)

	for i := 1; i <= number; i++ {
		name := fmt.Sprintf("%s%02d", prefix, i)
		_, err := gh.Repositories.Delete(ctx, org, name)
		if err != nil {
			return err
		}
	}

	return nil
}
