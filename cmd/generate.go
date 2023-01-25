/*
Copyright © 2023 SKIP, Kartverket <william.andersson@kartverket.no>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"

	skiperator "github.com/kartverket/skiperator/api/v1alpha1"
	"github.com/kartverket/skipper/pkg/util"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		// TODO Actually do something with errors

		application := skiperator.Application{}
		application.FillDefaultsSpec()
		application.FillDefaultsStatus()

		appName, err := promptStringWithRegexValidator("^[a-z]+(-[a-z]+)*$", "Application Name")
		if err != nil {
			return
		}

		appNamespace, err := promptStringWithRegexValidator("^[a-z]+(-[a-z]+)*$", "Application Namespace")
		if err != nil {
			return
		}

		application.Name = appName
		application.Namespace = appNamespace

		shouldWrite, err := promptYesNoSelect("Write Application to file?")
		if err != nil {
			return
		}

		if shouldWrite {
			filename, err := promptStringWithRegexValidator(`^\w+$`, "Filename (.yaml added automatically)")
			if err != nil {
				return
			}
			util.WriteApplicationToFile(filename, application)
		} else {
			appJson, _ := json.MarshalIndent(application, "", "\t")

			println(appJson)
		}

	},
}

func init() {
	skipAppCmd.AddCommand(generateCmd)
}

func promptStringWithRegexValidator(regex string, label string) (string, error) {
	validateWithRegex := func(input string) error {
		match, _ := regexp.MatchString(regex, input)
		if !match {
			errorMessage := fmt.Sprintf("%s must follow regex: %s", label, regex)
			err := errors.New(errorMessage)
			return err
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    label,
		Validate: validateWithRegex,
	}

	return prompt.Run()
}

func promptYesNoSelect(label string) (bool, error) {
	prompt := promptui.Select{
		Label: label,
		Items: []string{"Yes", "No"},
	}

	responseToBoolMap := map[string]bool{
		"Yes": true,
		"No":  false,
	}

	_, response, err := prompt.Run()
	if err != nil {
		return false, err
	}

	return responseToBoolMap[response], nil
}
