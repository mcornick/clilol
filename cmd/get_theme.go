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
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

type getThemeOutput struct {
	Response struct {
		Theme struct {
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
		} `json:"theme"`
	} `json:"response"`
	Request resultRequest `json:"request"`
}

var getThemeCmd = &cobra.Command{
	Use:   "theme [theme-name]",
	Short: "Get theme information",
	Long:  "Gets information about a theme.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		result, err := getTheme(args[0])
		cobra.CheckErr(err)
		if result.Request.Success {
			updatedAt, err := strconv.ParseInt(result.Response.Theme.Updated, 10, 64)
			cobra.CheckErr(err)
			fmt.Printf(
				"%s: %s by %s (%s) updated %s\n",
				result.Response.Theme.ID,
				result.Response.Theme.Name,
				result.Response.Theme.Author,
				result.Response.Theme.AuthorURL,
				time.Unix(updatedAt, 0),
			)
		} else {
			cobra.CheckErr(fmt.Errorf("%d", result.Request.StatusCode))
		}
	},
}

func init() {
	getCmd.AddCommand(getThemeCmd)
}

func getTheme(name string) (getThemeOutput, error) {
	var result getThemeOutput
	body := callAPIWithParams(
		http.MethodGet,
		"/theme/"+name+"/info",
		nil,
		true,
	)
	err := json.Unmarshal(body, &result)
	return result, err
}
