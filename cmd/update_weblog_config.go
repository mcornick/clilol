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
	updateWeblogConfigFilename string
	updateWeblogConfigCmd      = &cobra.Command{
		Use:   "config",
		Short: "set your weblog config",
		Long: `Sets configuration of your weblog.lol weblog.

The format is the same as obtained via "clilol weblog get config".

If you specify a filename with the --filename flag, the content of the file
will be used. If you do not specify a filename, the content will be read
from stdin.`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			type Result struct {
				Request  resultRequest `json:"request"`
				Response struct {
					Message string `json:"message"`
				} `json:"response"`
			}
			var result Result
			var content string
			if updateWeblogConfigFilename != "" {
				input, err := os.ReadFile(updateWeblogConfigFilename)
				cobra.CheckErr(err)
				content = string(input)
			} else {
				stdin, err := io.ReadAll(os.Stdin)
				cobra.CheckErr(err)
				content = string(stdin)
			}
			body := callAPIWithRawData(
				http.MethodPost,
				"/address/"+viper.GetString("address")+"/weblog/configuration",
				content,
				true,
			)
			err := json.Unmarshal(body, &result)
			cobra.CheckErr(err)
			if result.Request.Success {
				fmt.Println(result.Response.Message)
			} else {
				cobra.CheckErr(fmt.Errorf(result.Response.Message))
			}
		},
	}
)

func init() {
	updateWeblogConfigCmd.Flags().StringVarP(
		&updateWeblogConfigFilename,
		"filename",
		"f",
		"",
		"file to read config from (default stdin)",
	)
	updateWeblogCmd.AddCommand(updateWeblogConfigCmd)
}
