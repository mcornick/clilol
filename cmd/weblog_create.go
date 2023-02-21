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
	weblogCreateFilename string
	weblogCreateCmd      = &cobra.Command{
		Use:   "create",
		Short: "create a weblog entry",
		Long: `Creates an entry in your weblog.

If you specify a filename with the --filename flag, the content of the file
will be used. If you do not specify a filename, the content will be read
from stdin.`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			type Result struct {
				Request struct {
					StatusCode int  `json:"status_code"`
					Success    bool `json:"success"`
				} `json:"request"`
				Response struct {
					Message string `json:"message"`
					Entry   struct {
						Location string `json:"location"`
						Title    string `json:"title"`
						Date     int64  `json:"date"`
						Type     string `json:"type"`
						Status   string `json:"status"`
						Body     string `json:"body"`
						Source   string `json:"source"`
						Metadata struct {
							Date string `json:"date"`
							Slug string `json:"slug"`
						} `json:"metadata"`
						Output string `json:"output"`
						Entry  string `json:"entry"`
					} `json:"entry"`
				} `json:"response"`
			}
			var result Result
			var content string
			if weblogCreateFilename != "" {
				input, err := os.ReadFile(weblogCreateFilename)
				cobra.CheckErr(err)
				content = string(input)
			} else {
				stdin, err := io.ReadAll(os.Stdin)
				cobra.CheckErr(err)
				content = string(stdin)
			}
			body := callAPIWithRawData(
				http.MethodPost,
				"/address/"+viper.GetString("address")+"/weblog/entry",
				content,
				true,
			)
			err := json.Unmarshal(body, &result)
			cobra.CheckErr(err)
			if !silent {
				if !wantJson {
					if result.Request.Success {
						fmt.Println(result.Response.Message)
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
	weblogCreateCmd.Flags().StringVarP(
		&weblogCreateFilename,
		"filename",
		"f",
		"",
		"file to read entry from (default stdin)",
	)
	weblogCmd.AddCommand(weblogCreateCmd)
}
