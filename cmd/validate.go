package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	util "github.com/kartverket/skipper/pkg/util"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validates a .yaml file as an Application and returns the resulting Application as a file.",
	Long: `Validates a SKIP application by taking in the path to an Application in .yaml form.
			In addition the command writes the resulting Application to a new file, with
			added default values`,
	Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]

		application, err := util.ReadApplicationFromFile(filePath)

		if err != nil {
			fmt.Println("Your application was not formatted correctly. Error: %v" + err.Error())
			return
		}

		application.FillDefaultsSpec()
		application.FillDefaultsStatus()

		util.WriteApplicationToFile("test", application)

	},
}

func init() {
	skipAppCmd.AddCommand(validateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// validateCmd.PersistentFlags().String("foo", "", "A help for foo")
}
