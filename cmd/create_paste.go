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
	"io"
	"os"

	"github.com/ejstreet/omglol-client-go/omglol"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	createPasteFilename string
	createPasteListed   bool
	createPasteCmd      = &cobra.Command{
		Use:   "paste [title]",
		Short: "Create or update a paste",
		Long: `Create or update a paste in your pastebin.

If the specified title is already in use, that paste will be updated.
If the title is not in use, a new paste will be created.

If you specify a filename with the --filename flag, the content of the file
will be used. If you do not specify a filename, the content will be read
from stdin.

The paste will be created as unlisted by default. To create a listed
paste, use the --listed flag.`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			title := args[0]
			err := createPaste(title, createPasteFilename, createPasteListed)
			handleAPIError(err)
			fmt.Printf("Paste %s created.\n", title)
		},
	}
)

func init() {
	createPasteCmd.Flags().StringVarP(
		&createPasteFilename,
		"filename",
		"f",
		"",
		"file to read paste from (default stdin)",
	)
	createPasteCmd.Flags().BoolVarP(
		&createPasteListed,
		"listed",
		"l",
		false,
		"create paste as listed (default false)",
	)
	createCmd.AddCommand(createPasteCmd)
}

func createPaste(title string, filename string, listed bool) error {
	var content string
	if filename != "" {
		input, err := os.ReadFile(filename)
		cobra.CheckErr(err)
		content = string(input)
	} else {
		stdin, err := io.ReadAll(os.Stdin)
		cobra.CheckErr(err)
		content = string(stdin)
	}
	client, err := omglol.NewClient(viper.GetString("email"), viper.GetString("apikey"), endpoint)
	cobra.CheckErr(err)
	paste := omglol.NewPaste(title, content, listed)
	err = client.CreatePaste(viper.GetString("address"), *paste)
	return err
}
