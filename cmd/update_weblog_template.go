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

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	updateWeblogTemplateFilename string
	updateWeblogTemplateCmd      = &cobra.Command{
		Use:   "template",
		Short: "set your weblog template",
		Long: `Sets template of your weblog.lol weblog.

The format is the same as obtained via "clilol weblog get template".

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
				} `json:"response"`
			}
			var result Result
			var content string
			if updateWeblogTemplateFilename != "" {
				input, err := os.ReadFile(updateWeblogTemplateFilename)
				checkError(err)
				content = string(input)
			} else {
				stdin, err := io.ReadAll(os.Stdin)
				checkError(err)
				content = string(stdin)
			}
			body := callAPIWithRawData(
				http.MethodPost,
				"/address/"+viper.GetString("address")+"/weblog/template",
				content,
				true,
			)
			err := json.Unmarshal(body, &result)
			checkError(err)
			if result.Request.Success {
				log.Info(result.Response.Message)
			} else {
				checkError(fmt.Errorf(result.Response.Message))
			}
		},
	}
)

func init() {
	updateWeblogTemplateCmd.Flags().StringVarP(
		&updateWeblogTemplateFilename,
		"filename",
		"f",
		"",
		"file to read template from (default stdin)",
	)
	updateWeblogCmd.AddCommand(updateWeblogTemplateCmd)
}
