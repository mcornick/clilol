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
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	getPasteFilename string
	getPasteTitle    string
	getPasteCmd      = &cobra.Command{
		Use:   "paste",
		Short: "Get a paste",
		Long: `Gets a paste by title.

Specify the title with the --title flag.

The address can be specified with the --address flag. If not set,
it defaults to your own address.`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			type Result struct {
				Request  responseRequest `json:"request"`
				Response struct {
					Message string `json:"message"`
					Paste   struct {
						Title      string `json:"title"`
						Content    string `json:"content"`
						ModifiedOn int64  `json:"modified_on"`
					} `json:"paste"`
				} `json:"response"`
			}
			var result Result
			if addressFlag == "" {
				addressFlag = viper.GetString("address")
			}
			body := callAPIWithParams(
				http.MethodGet,
				"/address/"+addressFlag+"/pastebin/"+getPasteTitle,
				nil,
				true,
			)
			err := json.Unmarshal(body, &result)
			checkError(err)
			if result.Request.Success {
				if getPasteFilename != "" {
					err = os.WriteFile(getPasteFilename, []byte(result.Response.Paste.Content), 0o644)
					checkError(err)
				} else {
					fmt.Println(result.Response.Paste.Content)
				}
			} else {
				checkError(fmt.Errorf(result.Response.Message))
			}
		},
	}
)

func init() {
	getPasteCmd.Flags().StringVarP(
		&addressFlag,
		"address",
		"a",
		"",
		"address whose paste to get",
	)
	getPasteCmd.Flags().StringVarP(
		&getPasteTitle,
		"title",
		"t",
		"",
		"title of the paste to get",
	)
	getPasteCmd.Flags().StringVarP(
		&getPasteFilename,
		"filename",
		"f",
		"",
		"file to write paste to (default stdout)",
	)
	getPasteCmd.MarkFlagRequired("title")
	getCmd.AddCommand(getPasteCmd)
}
