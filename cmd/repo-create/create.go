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
	rootCmd.AddCommand(NewCreateCmd())
}

type createCmdOptions struct {
	IsPrivate bool
}

// NewCreateCmd generates the `create` command
func NewCreateCmd() *cobra.Command {
	s := createCmdOptions{}
	c := &cobra.Command{
		Use:   "create",
		Short: "Creates repositories",
		Long:  `Mass creates repositories`,
		RunE:  s.RunE,
	}
	c.Flags().BoolVar(&s.IsPrivate, "private", false, "Set repo to private")

	viper.BindPFlags(c.Flags())

	return c
}

func (s *createCmdOptions) RunE(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: authToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	gh := github.NewClient(tc)

	for i := 1; i <= number; i++ {
		name := fmt.Sprintf("%s%02d", prefix, i)
		_, _, err := gh.Repositories.Create(ctx, org, &github.Repository{
			Name:    &name,
			Private: &s.IsPrivate,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
