// SPDX-License-Identifier: MPL-2.0
//
// Copyright (c) 2023 Mark Cornick
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var manCmd = &cobra.Command{
	Use:    "man",
	Short:  "Generate man pages",
	Long:   fmt.Sprintln("Generates man pages for clilol in the current directory."),
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		header := &doc.GenManHeader{
			Title:   "Mark Cornick",
			Section: "1",
		}
		cobra.CheckErr(doc.GenManTree(rootCmd, header, "."))
	},
}

func init() {
	rootCmd.AddCommand(manCmd)
}
