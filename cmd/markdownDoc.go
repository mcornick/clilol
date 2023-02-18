// SPDX-License-Identifier: MIT
//
// Copyright (c) 2023 Mark Cornick
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use, copy,
// modify, merge, publish, distribute, sublicense, and/or sell copies
// of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS
// BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN
// ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var markdownDocCmd = &cobra.Command{
	Use:   "markdown-doc",
	Short: "Generate markdown docs",
	Long:  fmt.Sprintln("Generates markdown docs for clilol in the docs directory."),
	Run: func(cmd *cobra.Command, args []string) {
		// cobra.CheckErr(doc.GenMarkdownTree(rootCmd, "."))
		filePrepender := func(filename string) string {
			slug := strings.Replace(filepath.Base(filename), ".md", "", 1)
			title := strings.Replace(slug, "_", " ", -1)
			return fmt.Sprintf("---\ntitle: \"%s\"\n---\n", title)
		}
		linkHandler := func(name string) string {
			return name
		}
		_ = doc.GenMarkdownTreeCustom(rootCmd, "docs", filePrepender, linkHandler)
	},
}

func init() {
	rootCmd.AddCommand(markdownDocCmd)
}
