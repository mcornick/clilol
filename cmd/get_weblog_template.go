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
	getWeblogTemplateFilename string
	getWeblogTemplateCmd      = &cobra.Command{
		Use:   "template",
		Short: "Get your weblog template",
		Long: `Gets your weblog template in rendered form.

If you specify a filename with the --filename flag, the content will be written
to that file. If you do not specify a filename, the content will be written
to stdout.`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			type Result struct {
				Request  resultRequest `json:"request"`
				Response struct {
					Message  string `json:"message"`
					Template string `json:"template"`
				} `json:"response"`
			}
			var result Result
			body := callAPIWithParams(
				http.MethodGet,
				"/address/"+viper.GetString("address")+"/weblog/template",
				nil,
				true,
			)
			err := json.Unmarshal(body, &result)
			cobra.CheckErr(err)
			if result.Request.Success {
				if getWeblogTemplateFilename != "" {
					err = os.WriteFile(
						getWeblogTemplateFilename,
						[]byte(result.Response.Template),
						0o644,
					)
					cobra.CheckErr(err)
				} else {
					fmt.Println(result.Response.Template)
				}
			} else {
				cobra.CheckErr(fmt.Errorf(result.Response.Message))
			}
		},
	}
)

func init() {
	getWeblogTemplateCmd.Flags().StringVarP(
		&getWeblogTemplateFilename,
		"filename",
		"f",
		"",
		"file to write template to (default stdout)",
	)
	getWeblogCmd.AddCommand(getWeblogTemplateCmd)
}
