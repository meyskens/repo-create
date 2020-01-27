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
	rootCmd.AddCommand(NewRmLabelCmd())
}

type rmLabelCmdOptions struct {
	Prefix string
	Number int
	Org    string
	Name   string
}

// NewRmLabelCmd generates the `rmLabel` command
func NewRmLabelCmd() *cobra.Command {
	s := rmLabelCmdOptions{}
	c := &cobra.Command{
		Use:   "rm-label",
		Short: "remove label from repositories",
		Long:  `Mass removed a label from multiple repositories`,
		RunE:  s.RunE,
	}
	c.Flags().StringVarP(&s.Prefix, "prefix", "p", "", "Prefix of repository names")
	c.Flags().IntVarP(&s.Number, "number", "n", 1, "How many repositories to remove label from")
	c.Flags().StringVarP(&s.Org, "org", "o", "", "Org to remove label repos under")

	c.Flags().StringVarP(&s.Name, "name", "m", "", "Name of the label to remove")

	c.MarkFlagRequired("prefix")
	c.MarkFlagRequired("org")
	c.MarkFlagRequired("number")

	c.MarkFlagRequired("name")

	viper.BindPFlags(c.Flags())

	return c
}

func (s *rmLabelCmdOptions) RunE(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: authToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	gh := github.NewClient(tc)

	for i := 1; i <= s.Number; i++ {
		name := fmt.Sprintf("%s%02d", s.Prefix, i)
		_, err := gh.Issues.DeleteLabel(ctx, s.Org, name, s.Name)

		if err != nil {
			log.Println(err)
		}
	}

	return nil
}
