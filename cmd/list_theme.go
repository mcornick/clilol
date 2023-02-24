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

	"github.com/spf13/cobra"
)

type Theme struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Created       string `json:"created"`
	Updated       string `json:"updated"`
	Author        string `json:"author"`
	AuthorURL     string `json:"author_url"`
	Version       string `json:"version"`
	License       string `json:"license"`
	Description   string `json:"description"`
	PreviewCSS    string `json:"preview_css"`
	SampleProfile string `json:"sample_profile"`
	ThemeColor    string `json:"theme-color"`
}
type listThemeOutput struct {
	Request  resultRequest `json:"request"`
	Response struct {
		Message string           `json:"message"`
		Themes  map[string]Theme `json:"themes"`
	} `json:"response"`
}

var listThemeCmd = &cobra.Command{
	Use:   "theme",
	Short: "List profile themes",
	Long:  "Lists the available profile themes.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := listTheme()
		cobra.CheckErr(err)
		if result.Request.Success {
			fmt.Println(result.Response.Message)
			for _, theme := range result.Response.Themes {
				fmt.Printf("- %s\n", theme.ID)
			}
		} else {
			cobra.CheckErr(fmt.Errorf(result.Response.Message))
		}
	},
}

func init() {
	listCmd.AddCommand(listThemeCmd)
}

func listTheme() (listThemeOutput, error) {
	var result listThemeOutput
	body := callAPIWithParams(http.MethodGet, "/theme/list", nil, false)
	err := json.Unmarshal(body, &result)
	return result, err
}
