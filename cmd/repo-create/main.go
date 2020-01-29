package main

import (
	"flag"
	"os"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	authToken string
	prefix    string
	number    int
	start     int
	org       string

	rootCmd = &cobra.Command{
		Use:   "repo-create",
		Short: "repo-create is a tool for bulk creation of GitHub repositories",
		Long:  `repo-create is a tool for bulk creation of GitHub repositories`,
	}
)

func main() {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	viper.AutomaticEnv()

	rootCmd.PersistentFlags().StringVarP(&authToken, "auth-token", "t", os.Getenv("AUTH_TOKEN"), "GitHub auth token")
	rootCmd.PersistentFlags().StringVarP(&prefix, "prefix", "p", "", "Prefix of repository names")
	rootCmd.PersistentFlags().IntVarP(&number, "number", "n", 1, "How many repositories to create")
	rootCmd.PersistentFlags().IntVarP(&start, "start", "s", 1, "Number of the first repository")
	rootCmd.PersistentFlags().StringVarP(&org, "org", "o", "", "Org to create repos under")

	flag.Parse()
	err := rootCmd.Execute()
	if err != nil {
		glog.Error(err)
	}
}
