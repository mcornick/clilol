// SPDX-License-Identifier: MPL-2.0
//
// Copyright (c) 2023 Mark Cornick
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type createPasteInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Listed  int    `json:"listed,omitempty"`
}
type createPasteOutput struct {
	Request  resultRequest `json:"request"`
	Response struct {
		Message string `json:"message"`
		Title   string `json:"title"`
	} `json:"response"`
}

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
		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := createPaste(args[0], createPasteFilename, createPasteListed)
			if err != nil {
				return err
			}
			if result.Request.Success {
				fmt.Println(result.Response.Message)
			} else {
				return fmt.Errorf("%s", result.Response.Message)
			}
			return nil
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

func createPaste(title, filename string, listed bool) (createPasteOutput, error) {
	var result createPasteOutput
	var content string
	var listedInt int
	if filename != "" {
		input, err := os.ReadFile(filename) // #nosec G304
		if err != nil {
			return result, err
		}
		content = string(input)
	} else {
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			return result, err
		}
		content = string(stdin)
	}
	if listed {
		listedInt = 1
	} else {
		listedInt = 0
	}
	params := createPasteInput{title, content, listedInt}
	body, err := callAPIWithParams(
		http.MethodPost,
		"/address/"+viper.GetString("address")+"/pastebin",
		params,
		true,
	)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}
