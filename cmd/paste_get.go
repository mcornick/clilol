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
	pasteGetFilename string
	pasteGetTitle    string
	pasteGetCmd      = &cobra.Command{
		Use:   "get",
		Short: "get a paste",
		Long: `Gets a paste by title.

Specify the title with the --title flag.

The address can be specified with the --address flag. If not set,
it defaults to your own address.`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			type Result struct {
				Request struct {
					StatusCode int  `json:"status_code"`
					Success    bool `json:"success"`
				} `json:"request"`
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
			body := callAPIWithJSON(
				http.MethodGet,
				"/address/"+address+"/pastebin/"+pasteGetTitle,
				nil,
				true,
			)
			err := json.Unmarshal(body, &result)
			cobra.CheckErr(err)
			if !silent {
				if !wantJson {
					if result.Request.Success {
						if pasteGetFilename != "" {
							err = os.WriteFile(pasteGetFilename, []byte(result.Response.Paste.Content), 0o644)
							cobra.CheckErr(err)
						} else {
							fmt.Println(result.Response.Paste.Content)
						}
					} else {
						cobra.CheckErr(fmt.Errorf(result.Response.Message))
					}
				} else {
					fmt.Println(string(body))
				}
			}
		},
	}
)

func init() {
	pasteGetCmd.Flags().StringVarP(
		&address,
		"address",
		"a",
		viper.GetString("address"),
		"address whose paste to get",
	)
	pasteGetCmd.Flags().StringVarP(
		&pasteGetTitle,
		"title",
		"t",
		"",
		"title of the paste to get",
	)
	pasteGetCmd.Flags().StringVarP(
		&pasteGetFilename,
		"filename",
		"f",
		"",
		"file to write paste to (default stdout)",
	)
	pasteCmd.AddCommand(pasteGetCmd)
}
