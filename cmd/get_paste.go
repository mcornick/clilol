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
	getPasteCmd      = &cobra.Command{
		Use:   "paste [title]",
		Short: "Get a paste",
		Long: `Gets a paste by title.

The address can be specified with the --address flag. If not set,
it defaults to your own address.`,
		Args: cobra.ExactArgs(1),
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
				"/address/"+addressFlag+"/pastebin/"+args[0],
				nil,
				true,
			)
			err := json.Unmarshal(body, &result)
			cobra.CheckErr(err)
			if result.Request.Success {
				if getPasteFilename != "" {
					err = os.WriteFile(getPasteFilename, []byte(result.Response.Paste.Content), 0o644)
					cobra.CheckErr(err)
				} else {
					fmt.Println(result.Response.Paste.Content)
				}
			} else {
				cobra.CheckErr(fmt.Errorf(result.Response.Message))
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
		&getPasteFilename,
		"filename",
		"f",
		"",
		"file to write paste to (default stdout)",
	)
	getCmd.AddCommand(getPasteCmd)
}
