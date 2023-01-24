// /*
// Copyright © 2022 NAME HERE <EMAIL ADDRESS>
// */
package cmd

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	skiperatorv1alpha1 "github.com/kartverket/skiperator/api/v1alpha1"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// getAppsCmd represents the getApps command
var getAppsCmd = &cobra.Command{
	Use:   "get-apps",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var kubeconfig *string
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
		println(*kubeconfig)
		flag.Parse()

		// TODO Choose config
		config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			fmt.Println(err.Error())
		}

		cliSet, err := dynamic.NewForConfig(config)
		if err != nil {
			fmt.Println(err.Error())
		}

		gvr := schema.GroupVersionResource{
			Group:    "skiperator.kartverket.no",
			Version:  "v1alpha1",
			Resource: "applications",
		}

		// TODO Better namespace handling
		println("Fetching applications in current namespace: " + namespace)

		appList, err := cliSet.Resource(gvr).Namespace(namespace).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			fmt.Println(err)
		}

		// TODO Handle empty list
		for _, unspecifiedApplication := range appList.Items {
			application := skiperatorv1alpha1.Application{}

			err = runtime.DefaultUnstructuredConverter.FromUnstructured(unspecifiedApplication.UnstructuredContent(), &application)
			if err != nil {
				fmt.Printf("error %s, converting unstructured to application type", err.Error())
			}

			println("Application: " + application.GetName() + " | Status: " + application.Status.ApplicationStatus.Message)
		}

	},
}

func init() {
	rootCmd.AddCommand(getAppsCmd)
	// namespace = getAppsCmd.Flag("namespace").Value.String()
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getAppsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getAppsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
