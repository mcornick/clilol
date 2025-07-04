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

type getWeblogConfigOutput struct {
	Response struct {
		Message       string `json:"message"`
		Configuration struct {
			Object struct {
				WeblogTitle                 string `json:"weblog-title"`
				WeblogDescription           string `json:"weblog-description"`
				Author                      string `json:"author"`
				Separator                   string `json:"separator"`
				TagPath                     string `json:"tag-path"`
				Timezone                    string `json:"timezone"`
				DateFormat                  string `json:"date-format"`
				DefaultPost                 string `json:"default-post"`
				FeedPostCount               string `json:"feed-post-count"`
				PostPathFormat              string `json:"post-path-format"`
				RecentPostsCount            string `json:"recent-posts-count"`
				RecentPostsFormat           string `json:"recent-posts-format"`
				PostListFormat              string `json:"post-list-format"`
				SearchStatus                string `json:"search-status"`
				SearchoutputsSuccessMessage string `json:"search-results-success-message"`
				SearchoutputsFailureMessage string `json:"search-results-failure-message"`
				SearchoutputsFormat         string `json:"search-results-format"`
			} `json:"object"`
			JSON string `json:"json"`
			Raw  string `json:"raw"`
		} `json:"configuration"`
	} `json:"response"`
	Request resultRequest `json:"request"`
}

var (
	getWeblogConfigFilename string
	getWeblogConfigCmd      = &cobra.Command{
		Use:   "config",
		Short: "Get your weblog config",
		Long: `Gets your weblog configuration in editable form.

If you specify a filename with the --filename flag, the content will be written
to that file. If you do not specify a filename, the content will be written
to stdout.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := getWeblogConfig()
			if err != nil {
				return err
			}
			if result.Request.Success {
				if getWeblogConfigFilename != "" {
					err = os.WriteFile(
						getWeblogConfigFilename,
						[]byte(result.Response.Configuration.Raw),
						0o600,
					)
					if err != nil {
						return err
					}
				} else {
					fmt.Println(result.Response.Configuration.Raw)
				}
			} else {
				return fmt.Errorf("%s", result.Response.Message)
			}
			return nil
		},
	}
)

func init() {
	getWeblogConfigCmd.Flags().StringVarP(
		&getWeblogConfigFilename,
		"filename",
		"f",
		"",
		"file to write configuration to (default stdout)",
	)
	getWeblogCmd.AddCommand(getWeblogConfigCmd)
}

func getWeblogConfig() (getWeblogConfigOutput, error) {
	var result getWeblogConfigOutput
	body, err := callAPIWithParams(
		http.MethodGet,
		"/address/"+viper.GetString("address")+"/weblog/configuration",
		nil,
		true,
	)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}
