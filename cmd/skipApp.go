package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// skipAppCmd represents the skipApp command
var skipAppCmd = &cobra.Command{
	Use:     "skip-app",
	Example: "skip-app list --namespace a-namespace",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("It seems you did not give any arguments to skip-app. Please see skipper skip-app --help for more")
	},
}

func init() {
	rootCmd.AddCommand(skipAppCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// skipAppCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// skipAppCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
