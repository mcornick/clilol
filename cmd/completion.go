// SPDX-License-Identifier: MPL-2.0
//
// Copyright (c) 2023 Mark Cornick
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:    "completion",
	Short:  "generate completion script",
	Long:   "Generates completion scripts for shells.",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cobra.CheckErr(cmd.Root().GenBashCompletion(os.Stdout))
		case "zsh":
			cobra.CheckErr(cmd.Root().GenZshCompletion(os.Stdout))
		case "fish":
			cobra.CheckErr(cmd.Root().GenFishCompletion(os.Stdout, true))
		case "powershell":
			cobra.CheckErr(cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout))
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
