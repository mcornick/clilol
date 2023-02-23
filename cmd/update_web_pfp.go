// SPDX-License-Identifier: MPL-2.0
//
// Copyright (c) 2023 Mark Cornick
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	updateWebPFPFilename string
	updateWebPFPCmd      = &cobra.Command{
		Use:   "pfp",
		Short: "set your profile picture",
		Long: `Sets your profile picture.

Specify an image file with the --filename flag.`,
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
			content, err := os.ReadFile(updateWebPFPFilename)
			checkError(err)
			encoded := "data:text/plain;base64," + base64.StdEncoding.EncodeToString(content)
			body := callAPIWithRawData(
				http.MethodPost,
				"/address/"+viper.GetString("address")+"/pfp",
				encoded,
				true,
			)
			err = json.Unmarshal(body, &result)
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
	updateWebPFPCmd.Flags().StringVarP(
		&updateWebPFPFilename,
		"filename",
		"f",
		"",
		"file to read PFP from (default stdin)",
	)
	updateWebCmd.AddCommand(updateWebPFPCmd)
}
