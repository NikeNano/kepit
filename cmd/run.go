/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"kepit/cmd/utils"
	"os"
	"plugin"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := cmd.Flags().GetString("file")
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		// Create a tmp file for the new file to live in
		tmpFile, err := ioutil.TempFile("tmp", "*Temporary.go")
		if err != nil {
			fmt.Println(errors.New("Error creating tmp file for new golang file"))
			os.Exit(1)
		}
		// We have to clean up ourselfs after words
		// [TODO] make a function to do all the clean up
		defer os.Remove(tmpFile.Name())

		err = utils.ParseFile(file, tmpFile.Name())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Need a temporary file for the plugin as well
		tmpPlugin, err := ioutil.TempFile("tmp", "*Temporary.so")
		if err != nil {
			fmt.Println(errors.New("Error creating tmp for plugin file"))
			os.Exit(1)
		}
		defer os.Remove(tmpFile.Name())
		err = utils.MakePlugin(tmpFile.Name(), tmpPlugin.Name())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Load the plugin here

		plug, err := plugin.Open(tmpPlugin.Name())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		structNameArg, err := cmd.Flags().GetString("structName")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		structName, err := plug.Lookup(structNameArg)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Read the data to the struct of interest
		jsonFile, err := cmd.Flags().GetString("jsonFile")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		inputJSON, err := ioutil.ReadFile(jsonFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = json.Unmarshal([]byte(inputJSON), &structName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Marshal the file back
		oputputJSON, err := json.Marshal(structName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		outputPath, err := cmd.Flags().GetString("outputPath")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = ioutil.WriteFile(outputPath, oputputJSON, 0644)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Your results can be found here: ", outputPath)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringP("file", "f", "", "File to parse for structs")
	runCmd.Flags().StringP("jsonFile", "j", "", "The input json file that should be parsed")
	runCmd.Flags().StringP("structName", "s", "", "The name of the struct which we should parse into")
	runCmd.Flags().StringP("outputPath", "o", "", "The path for the output file")

	runCmd.MarkFlagRequired("file")
	runCmd.MarkFlagRequired("structName")
	runCmd.MarkFlagRequired("jsonFile")
	runCmd.MarkFlagRequired("outputPath")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func readFile(args []string) {

}
