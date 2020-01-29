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
	Name  string
	Color string
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
	c.Flags().StringVarP(&s.Name, "name", "m", "", "Name of the label")
	c.Flags().StringVarP(&s.Color, "color", "c", "", "Color for the label")

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

	for i := 1; i <= number; i++ {
		name := fmt.Sprintf("%s%02d", prefix, i)
		_, _, err := gh.Issues.CreateLabel(ctx, org, name, &github.Label{
			Name:  &s.Name,
			Color: &s.Color,
		})

		if err != nil {
			log.Println(err)
		}
	}

	return nil
}
