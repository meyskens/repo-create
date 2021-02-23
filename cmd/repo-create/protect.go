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
	rootCmd.AddCommand(NewProtectCmd())
}

type protectCmdOptions struct {
	Branch       string
	MinReviewers int
}

// NewProtectCmd generates the `protect` command
func NewProtectCmd() *cobra.Command {
	s := protectCmdOptions{}
	c := &cobra.Command{
		Use:   "protect",
		Short: "add protect to repositories",
		Long:  `Mass add a protect to multiple repositories`,
		RunE:  s.RunE,
	}

	c.Flags().StringVarP(&s.Branch, "branch", "b", "master", "Name of the branch")
	c.Flags().IntVar(&s.MinReviewers, "min-reviewers", 1, "Minimal reviewers")

	c.MarkFlagRequired("prefix")
	c.MarkFlagRequired("org")
	c.MarkFlagRequired("number")

	viper.BindPFlags(c.Flags())

	return c
}

func (s *protectCmdOptions) RunE(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: authToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	gh := github.NewClient(tc)

	for i := start; i <= number; i++ {
		name := fmt.Sprintf("%s%02d", prefix, i)
		_, _, err := gh.Repositories.UpdateBranchProtection(ctx, org, name, s.Branch, &github.ProtectionRequest{
			RequiredPullRequestReviews: &github.PullRequestReviewsEnforcementRequest{RequiredApprovingReviewCount: s.MinReviewers},
		})

		if err != nil {
			log.Println(err)
		}
	}

	return nil
}
