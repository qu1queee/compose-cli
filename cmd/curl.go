// Copyright © 2018 Enrique Encalada
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/HeavyWombat/dyff/pkg/v1/bunt"
	"github.com/qu1queee/compose-cli/pkg/compose_helper"
	"github.com/spf13/cobra"
)

type argError struct {
	arg  string
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%s %s - %s", bunt.Colorize("ERROR:", bunt.Red), bunt.Colorize(e.arg, bunt.FireBrick), bunt.Colorize(e.prob, bunt.LightCoral))
}

//PathSanityChecks checks for curl path argument
func PathSanityChecks(path string) {
	switch {
	case !strings.ContainsAny(path, "/"):
		fmt.Printf("%v", &argError{"curl " + path, "missing / \n"})
		os.Exit(3)
	case strings.ContainsAny(path, " "):
		fmt.Printf("%v", &argError{"curl " + path, "Should not used whitespaces \n"})
		os.Exit(3)
	}
}

// curlCmd represents the curl command
var curlCmd = &cobra.Command{
	Use:   "curl /api-call-path",
	Args:  cobra.ExactArgs(1),
	Short: "Trigger an specific API call",
	Long: `curl will trigger a GET http method agains the API call type, for more
information visit:
https://www.compose.com/articles/the-ibm-cloud-compose-api/.

Example of API calls

Retrieve current Composedb status:
compose-cli --deployment-id some-id --foundation-endpoint some-endpoint --api-token some-token curl /alerts

Retrieve available logfiles:
compose-cli --deployment-id some-id --foundation-endpoint some-endpoint --api-token some-token curl /logfiles

`,
	Run: func(cmd *cobra.Command, args []string) {
		PathSanityChecks(args[0])
		result, err := compose_helper.Curl(deploymentID, foundationEndpoint, apiToken, args[0])
		if err != nil {
			fmt.Printf(err.Error())
		}
		fmt.Print(bunt.Colorize(result, bunt.LimeGreen) + "\n")
	},
}

func init() {
	rootCmd.AddCommand(curlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// curlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// curlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
