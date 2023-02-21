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
	nowSetFilename string
	nowSetListed   bool
	nowSetCmd      = &cobra.Command{
		Use:   "set",
		Short: "set Now page content",
		Long: `Sets your Now page content.

If you specify a filename with the --filename flag, the content of the file
will be used. If you do not specify a filename, the content will be read
from stdin.`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			type Input struct {
				Content string `json:"content"`
				Listed  int    `json:"listed"`
			}
			type Result struct {
				Request struct {
					StatusCode int  `json:"status_code"`
					Success    bool `json:"success"`
				} `json:"request"`
				Response struct {
					Message string `json:"message"`
				} `json:"response"`
			}
			var result Result
			var listed int
			var content string
			if nowSetFilename != "" {
				input, err := os.ReadFile(nowSetFilename)
				cobra.CheckErr(err)
				content = string(input)
			} else {
				stdin, err := io.ReadAll(os.Stdin)
				cobra.CheckErr(err)
				content = string(stdin)
			}
			if nowSetListed {
				listed = 1
			} else {
				listed = 0
			}
			nowPage := Input{content, listed}
			body := callAPIWithJSON(
				http.MethodPost,
				"/address/"+viper.GetString("address")+"/now",
				nowPage,
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
	nowSetCmd.Flags().StringVarP(
		&nowSetFilename,
		"filename",
		"f",
		"",
		"file to read Now page from (default stdin)",
	)
	nowSetCmd.Flags().BoolVarP(
		&nowSetListed,
		"listed",
		"l",
		false,
		"create Now page as listed (default false)",
	)
	nowCmd.AddCommand(nowSetCmd)
}
