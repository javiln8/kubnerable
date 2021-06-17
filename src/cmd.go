package main

import (
	"github.com/spf13/cobra"
)

// Flag variables
var (
	vulsFilePathFlag string
)

var rootCmd = &cobra.Command{
	Use:     "kubnerable",
	Short:   "Scan a Kubernetes cluster for misconfigurations and vulnerabilities",
	Long:    "Kubnerable scans all pods and containers of a Kubernetes cluster in search of exploitable vulnerabilities",
	Example: "kubnerable -f ../vulnerabilities.yaml",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		AnalyzePods(vulsFilePathFlag)
	},
}

// ExecuteCmd is the main Cobra command
func ExecuteCmd() {
	err := rootCmd.Execute()
	CheckError(err, "Could not start the CMD")
}

// init initialises available commands and flags
// https://github.com/spf13/cobra#create-rootcmd
func init() {
	// Flags
	rootCmd.PersistentFlags().StringVarP(&vulsFilePathFlag, "vuls-file", "f", "../vulnerabilities.yaml", "Vulnerabilities YAML file path, needed for scanning")
}
