package cmd

import (
	"context"
	"fmt"

	"github.com/coreos/go-oidc"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2/google"
	"k8s.io/client-go/tools/clientcmd"
)

func getProjectFromKubernetesConfig() string {
	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	).RawConfig()

	if err != nil {
		fmt.Println("Something went wrong")
		fmt.Println(err.Error())
	}

	return config.Contexts[config.CurrentContext].Cluster
}

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Shows the current project and email",
	Long: `Fetches the current project and email using default credentials from Google.
	See https://cloud.google.com/docs/authentication/application-default-credentials`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		creds, err := google.FindDefaultCredentials(ctx)
		if err != nil {
			fmt.Println("Default credentials were not found. Try logging in to google cloud using the following command:")
			fmt.Println("gcloud auth application-default login")
		}

		provider, err := oidc.NewProvider(ctx, "https://accounts.google.com")
		if err != nil {
			fmt.Println("Something went wrong")
			fmt.Println(err.Error())
		}

		user, err := provider.UserInfo(ctx, creds.TokenSource)
		if err != nil {
			fmt.Println("Could not find information regarding the user. Are you sure your application defaults are correct?")
		}

		fmt.Println("Current email: " + user.Email)
		fmt.Println("Current project/context: " + getProjectFromKubernetesConfig())
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
}
