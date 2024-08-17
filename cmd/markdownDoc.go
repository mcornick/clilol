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
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var markdownDocCmd = &cobra.Command{
	Use:    "markdown-doc",
	Short:  "Generate markdown docs",
	Long:   fmt.Sprintln("Generates markdown docs for clilol in the docs directory."),
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		filePrepender := func(filename string) string {
			slug := strings.Replace(filepath.Base(filename), ".md", "", 1)
			title := strings.Replace(slug, "_", " ", -1)
			return fmt.Sprintf("---\ntitle: \"%s\"\n---\n", title)
		}
		linkHandler := func(name string) string {
			return name
		}
		err := doc.GenMarkdownTreeCustom(rootCmd, "docs/commands", filePrepender, linkHandler)
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(markdownDocCmd)
}
