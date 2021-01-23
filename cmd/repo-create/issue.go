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
	rootCmd.AddCommand(NewIssueCmd())
}

type issueCmdOptions struct {
	Title   string
	Content string
	Labels  []string
}

// NewIssueCmd generates the `label` command
func NewIssueCmd() *cobra.Command {
	s := issueCmdOptions{}
	c := &cobra.Command{
		Use:   "issue",
		Short: "add issue to repositories",
		Long:  `Mass add a issue to multiple repositories`,
		RunE:  s.RunE,
	}
	c.Flags().StringVarP(&s.Title, "title", "m", "", "Title of the issue")
	c.Flags().StringVarP(&s.Content, "content", "c", "", "Content of the issue")
	c.Flags().StringArrayVar(&s.Labels, "labels", []string{}, "Labels for the issue")

	c.MarkFlagRequired("title")
	c.MarkFlagRequired("content")

	viper.BindPFlags(c.Flags())

	return c
}

func (i *issueCmdOptions) RunE(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: authToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	gh := github.NewClient(tc)

	for j := start; j <= number; j++ {
		name := fmt.Sprintf("%s%02d", prefix, j)
		_, _, err := gh.Issues.Create(ctx, org, name, &github.IssueRequest{
			Title:  &i.Title,
			Body:   &i.Content,
			Labels: &i.Labels,
		})

		if err != nil {
			log.Println(err)
		}
	}

	return nil
}
