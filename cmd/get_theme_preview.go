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

type getThemePreviewOutput struct {
	Response struct {
		Message string `json:"message"`
		HTML    string `json:"html"`
	} `json:"response"`
}

var (
	getThemePreviewFilename string
	getThemePreviewCmd      = &cobra.Command{
		Use:   "preview [theme-name]",
		Short: "Get theme preview",
		Long: `Gets an HTML preview of a theme.

If you specify a filename with the --filename flag, the content will be written
to that file. If you do not specify a filename, the content will be written
to stdout.`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			validateConfig()
			result, err := getThemePreview(args[0])
			cobra.CheckErr(err)
			if getThemePreviewFilename != "" {
				err = os.WriteFile(getThemePreviewFilename, []byte(result.Response.HTML), 0o644)
				cobra.CheckErr(err)
			} else {
				fmt.Println(result.Response.HTML)
			}
		},
	}
)

func init() {
	getThemePreviewCmd.Flags().StringVarP(
		&getThemePreviewFilename,
		"filename",
		"f",
		"",
		"file to write preview to (default stdout)",
	)
	getThemeCmd.AddCommand(getThemePreviewCmd)
}

func getThemePreview(name string) (getThemePreviewOutput, error) {
	var result getThemePreviewOutput
	body := callAPIWithParams(
		http.MethodGet,
		"/theme/"+name+"/preview",
		nil,
		true,
	)
	err := json.Unmarshal(body, &result)
	return result, err
}
