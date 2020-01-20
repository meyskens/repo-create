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

	flag.Parse()
	err := rootCmd.Execute()
	if err != nil {
		glog.Error(err)
	}
}
