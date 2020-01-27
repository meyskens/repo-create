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
	rootCmd.AddCommand(NewLabelCmd())
}

type labelCmdOptions struct {
	Prefix string
	Number int
	Org    string
	Name   string
	Color  string
}

// NewLabelCmd generates the `label` command
func NewLabelCmd() *cobra.Command {
	s := labelCmdOptions{}
	c := &cobra.Command{
		Use:   "label",
		Short: "add label to repositories",
		Long:  `Mass add a label to multiple repositories`,
		RunE:  s.RunE,
	}
	c.Flags().StringVarP(&s.Prefix, "prefix", "p", "", "Prefix of repository names")
	c.Flags().IntVarP(&s.Number, "number", "n", 1, "How many repositories to label")
	c.Flags().StringVarP(&s.Org, "org", "o", "", "Org to label repos under")

	c.Flags().StringVarP(&s.Name, "name", "m", "", "Name of the label")
	c.Flags().StringVarP(&s.Color, "color", "c", "", "Color for the label")

	c.MarkFlagRequired("prefix")
	c.MarkFlagRequired("org")
	c.MarkFlagRequired("number")

	c.MarkFlagRequired("name")
	c.MarkFlagRequired("color")

	viper.BindPFlags(c.Flags())

	return c
}

func (s *labelCmdOptions) RunE(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: authToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	gh := github.NewClient(tc)

	for i := 1; i <= s.Number; i++ {
		name := fmt.Sprintf("%s%02d", s.Prefix, i)
		_, _, err := gh.Issues.CreateLabel(ctx, s.Org, name, &github.Label{
			Name:  &s.Name,
			Color: &s.Color,
		})

		if err != nil {
			log.Println(err)
		}
	}

	return nil
}
