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
	"fmt"
	"os"
	"strings"

	"github.com/jordanbecketmoore/oapirouter/pkg/oapirouter"
	"github.com/pb33f/libopenapi"
	"github.com/spf13/cobra"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
	"sigs.k8s.io/yaml"
)

var name string
var namespace string
var gatewayName string
var inputFile string
var outputFile string
var hostnames string

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
	oapirouter --httproute-name my-http-route --gateway-name my-gateway --input /path/to/openapi.yaml

This command will generate an HTTPRoute resource based on the provided OpenAPI spec and print it to stdout.`,

	Run: func(cmd *cobra.Command, args []string) {
		// Load input file
		if inputFile == "" {
			cmd.PrintErrln("Error: --input flag is required") // TODO enable reading from stdin
			os.Exit(1)
		}

		inputData, err := os.ReadFile(inputFile)
		if err != nil {
			cmd.PrintErrf("Error reading input file: %v\n", err)
			os.Exit(1)
		}
		document, err := libopenapi.NewDocument(inputData)
		if err != nil {
			fmt.Errorf("Failed to create document: %v", err)
		}

		// Build document model
		documentModel, errs := document.BuildV3Model()
		if len(errs) > 0 {
			fmt.Errorf("Expected no errors, but got %v", errs)
		}
		httpRoute, err := oapirouter.DocumentModelToHTTPRoute(documentModel)
		if err != nil {
			fmt.Errorf("Expected no error, but got %v", err)
		}

		// Set hostnames for HTTPRoute
		if hostnames != "" {
			// Split hostnames by comma and add to HTTPRoute
			hostnameList := strings.Split(hostnames, ",")
			for _, hostname := range hostnameList {
				httpRoute.Spec.Hostnames = append(httpRoute.Spec.Hostnames, gatewayv1.Hostname(hostname))
			}
		}

		// Marhshal the HTTPRoute to YAML for better readability
		httpRouteYAML, err := yaml.Marshal(httpRoute)
		if err != nil {
			fmt.Errorf("Failed to marshal HTTPRoute to YAML: %v", err)
		}

		if outputFile != "" {
			err = os.WriteFile(outputFile, httpRouteYAML, 0644)
			if err != nil {
				cmd.PrintErrf("Error writing output file: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("HTTPRoute resource written to %s\n", outputFile)
		} else {
			fmt.Printf("%s", string(httpRouteYAML))
		}

	},
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
	rootCmd.PersistentFlags().StringVar(&name, "httproute-name", "", "Specify a name for the HTTPRoute resource")
	rootCmd.PersistentFlags().StringVar(&namespace, "namespace", "", "Specify a namespace for the HTTPRoute resource")
	rootCmd.PersistentFlags().StringVar(&gatewayName, "gateway-name", "", "Specify a gateway for the HTTPRoute resource")
	rootCmd.PersistentFlags().StringVar(&inputFile, "input", "", "Location of the OpenAPI spec file")
	rootCmd.PersistentFlags().StringVar(&outputFile, "output", "", "Location of the OpenAPI spec file")
	rootCmd.PersistentFlags().StringVar(&hostnames, "hostnames", "", "Specify hostnames as a comma-separated list for the HTTPRoute resource")
}
