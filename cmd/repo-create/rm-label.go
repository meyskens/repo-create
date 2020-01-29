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
	Name string
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

	c.Flags().StringVarP(&s.Name, "name", "m", "", "Name of the label to remove")

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

	for i := start; i <= number; i++ {
		name := fmt.Sprintf("%s%02d", prefix, i)
		_, err := gh.Issues.DeleteLabel(ctx, org, name, s.Name)

		if err != nil {
			log.Println(err)
		}
	}

	return nil
}
