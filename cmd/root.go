/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

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
	"os"

	"github.com/spf13/cobra"
)

var name string
var namespace string
var gatewayName string
var inputFile string
var outputFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "oapirouter",
	Short: "A CLI tool to generate HTTPRoute resources from OpenAPI specs",
	Long: `oapirouter is a command-line tool that helps you generate
HTTPRoute resources for Kubernetes from OpenAPI specifications.
It simplifies the process of creating and managing HTTPRoute resources
to handle the routing of HTTP requests in your applications inside a Kubernetes cluster
using Gateway API resources.

Example usage:
oapirouter --name my-route --namespace my-namespace --gateway-name my-gateway --input /path/to/openapi.yaml
This command will generate an HTTPRoute resource based on the provided OpenAPI spec and print it to stdout.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.oapirouter.yaml)")
	rootCmd.PersistentFlags().StringVar(&name, "name", "", "Specify a name for the HTTPRoute resource")
	rootCmd.PersistentFlags().StringVar(&namespace, "namespace", "", "Specify a namespace for the HTTPRoute resource")
	rootCmd.PersistentFlags().StringVar(&gatewayName, "gateway-name", "", "Specify a gateway for the HTTPRoute resource")
	rootCmd.PersistentFlags().StringVar(&inputFile, "input", "", "Location of the OpenAPI spec file")
	rootCmd.PersistentFlags().StringVar(&outputFile, "output", "", "Location of the OpenAPI spec file")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
