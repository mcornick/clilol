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
)

var (
	getThemePreviewFilename string
	getThemePreviewCmd      = &cobra.Command{
		Use:   "preview",
		Short: "get theme preview",
		Long: `Gets an HTML preview of a theme.

Specify the theme name with the --name flag.

If you specify a filename with the --filename flag, the content will be written
to that file. If you do not specify a filename, the content will be written
to stdout.`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			type Result struct {
				Response struct {
					Message string `json:"message"`
					HTML    string `json:"html"`
				} `json:"response"`
			}
			var result Result
			body := callAPIWithJSON(
				http.MethodGet,
				"/theme/"+nameFlag+"/preview",
				nil,
				true,
			)
			err := json.Unmarshal(body, &result)
			cobra.CheckErr(err)
			if !silent {
				if !wantJson {
					if getThemePreviewFilename != "" {
						err = os.WriteFile(getThemePreviewFilename, []byte(result.Response.HTML), 0o644)
						cobra.CheckErr(err)
					} else {
						fmt.Println(result.Response.HTML)
					}
				} else {
					fmt.Println(string(body))
				}
			}
		},
	}
)

func init() {
	getThemePreviewCmd.Flags().StringVarP(
		&nameFlag,
		"name",
		"n",
		"",
		"name of the theme",
	)
	getThemePreviewCmd.Flags().StringVarP(
		&getThemePreviewFilename,
		"filename",
		"f",
		"",
		"file to write preview to (default stdout)",
	)
	getThemeCmd.AddCommand(getThemePreviewCmd)
}
