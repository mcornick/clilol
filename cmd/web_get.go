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
	webGetFilename string
	webGetCmd      = &cobra.Command{
		Use:   "get",
		Short: "get your webpage content",
		Long: `Gets the Markdown content for your webpage.

If you specify a filename with the --filename flag, the content will be written
to that file. If you do not specify a filename, the content will be written
to stdout.`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			type Result struct {
				Request struct {
					StatusCode int  `json:"status_code"`
					Success    bool `json:"success"`
				} `json:"request"`
				Response struct {
					Message  string `json:"message"`
					Content  string `json:"content"`
					Type     string `json:"type"`
					Theme    string `json:"theme"`
					CSS      string `json:"css"`
					Head     string `json:"head"`
					Verified int    `json:"verified"`
					PFP      string `json:"pfp"`
					Metadata string `json:"metadata"`
					Branding string `json:"branding"`
					Modified string `json:"modified"`
				} `json:"response"`
			}
			var result Result
			body := callAPI(
				http.MethodGet,
				"/address/"+viper.GetString("address")+"/web",
				nil,
				true,
			)
			err := json.Unmarshal(body, &result)
			cobra.CheckErr(err)
			if !silent {
				if !wantJson {
					if result.Request.Success {
						if webGetFilename != "" {
							err = os.WriteFile(webGetFilename, []byte(result.Response.Content), 0o644)
							cobra.CheckErr(err)
						} else {
							fmt.Println(result.Response.Content)
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
	webGetCmd.Flags().StringVarP(
		&webGetFilename,
		"filename",
		"f",
		"",
		"file to write webpage to (default stdout)",
	)
	webCmd.AddCommand(webGetCmd)
}
