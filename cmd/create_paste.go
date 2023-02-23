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

var (
	createPasteTitle    string
	createPasteFilename string
	createPasteListed   bool
	createPasteCmd      = &cobra.Command{
		Use:   "paste",
		Short: "Create or update a paste",
		Long: `Create or update a paste in your pastebin.

Specify a title with the --title flag. If the title is already in use,
that paste will be updated. If the title is not in use, a new paste will
be created.

If you specify a filename with the --filename flag, the content of the file
will be used. If you do not specify a filename, the content will be read
from stdin.`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			type Input struct {
				Title   string `json:"title"`
				Content string `json:"content"`
				Listed  int    `json:"listed,omitempty"`
			}
			type Result struct {
				Request struct {
					StatusCode int  `json:"status_code"`
					Success    bool `json:"success"`
				} `json:"request"`
				Response struct {
					Message string `json:"message"`
					Title   string `json:"title"`
				} `json:"response"`
			}
			var result Result
			var listed int
			var content string
			if createPasteFilename != "" {
				input, err := os.ReadFile(createPasteFilename)
				checkError(err)
				content = string(input)
			} else {
				stdin, err := io.ReadAll(os.Stdin)
				checkError(err)
				content = string(stdin)
			}
			if createPasteListed {
				listed = 1
			} else {
				listed = 0
			}
			paste := Input{createPasteTitle, content, listed}
			body := callAPIWithParams(
				http.MethodPost,
				"/address/"+viper.GetString("address")+"/pastebin",
				paste,
				true,
			)
			err := json.Unmarshal(body, &result)
			checkError(err)
			if !jsonFlag {
				if result.Request.Success {
					logInfo(result.Response.Message)
				} else {
					checkError(fmt.Errorf(result.Response.Message))
				}
			} else {
				fmt.Println(string(body))
			}
		},
	}
)

func init() {
	createPasteCmd.Flags().StringVarP(
		&createPasteTitle,
		"title",
		"t",
		"",
		"title of the paste to create",
	)
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
