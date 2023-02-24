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
	updateNowFilename string
	updateNowListed   bool
	updateNowCmd      = &cobra.Command{
		Use:   "set",
		Short: "set Now page content",
		Long: `Sets your Now page content.

If you specify a filename with the --filename flag, the content of the file
will be used. If you do not specify a filename, the content will be read
from stdin.

The Now page will be created as unlisted by default. To create a listed
Now page, use the --listed flag.`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			type Input struct {
				Content string `json:"content"`
				Listed  int    `json:"listed"`
			}
			type Result struct {
				Request  responseRequest `json:"request"`
				Response struct {
					Message string `json:"message"`
				} `json:"response"`
			}
			var result Result
			var listed int
			var content string
			if updateNowFilename != "" {
				input, err := os.ReadFile(updateNowFilename)
				checkError(err)
				content = string(input)
			} else {
				stdin, err := io.ReadAll(os.Stdin)
				checkError(err)
				content = string(stdin)
			}
			if updateNowListed {
				listed = 1
			} else {
				listed = 0
			}
			nowPage := Input{content, listed}
			body := callAPIWithParams(
				http.MethodPost,
				"/address/"+viper.GetString("address")+"/now",
				nowPage,
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
	updateNowCmd.Flags().StringVarP(
		&updateNowFilename,
		"filename",
		"f",
		"",
		"file to read Now page from (default stdin)",
	)
	updateNowCmd.Flags().BoolVarP(
		&updateNowListed,
		"listed",
		"l",
		false,
		"create Now page as listed (default false)",
	)
	updateCmd.AddCommand(updateNowCmd)
}
