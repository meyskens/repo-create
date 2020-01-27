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
	Prefix    string
	Number    int
	Org       string
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
	c.Flags().StringVarP(&s.Prefix, "prefix", "p", "", "Prefix of repository names")
	c.Flags().IntVarP(&s.Number, "number", "n", 1, "How many repositories to create")
	c.Flags().StringVarP(&s.Org, "org", "o", "", "Org to create repos under")
	c.Flags().BoolVar(&s.IsPrivate, "private", false, "Set repo to private")

	c.MarkFlagRequired("prefix")
	c.MarkFlagRequired("org")
	c.MarkFlagRequired("number")

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

	for i := 1; i <= s.Number; i++ {
		name := fmt.Sprintf("%s%02d", s.Prefix, i)
		_, _, err := gh.Repositories.Create(ctx, s.Org, &github.Repository{
			Name:    &name,
			Private: &s.IsPrivate,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
