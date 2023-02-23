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
	"github.com/spf13/viper"
)

var (
	createStatusEmoji            string
	createStatusStatus           string
	createStatusSkipMastodonPost bool
	createStatusCmd              = &cobra.Command{
		Use:   "status",
		Short: "Create a status",
		Long: `Posts a status to status.lol.

Specify the status text with the --text flag.
Quote the text if it contains spaces.

You can specify an emoji with the --emoji flag. This must be an
actual emoji, not a :emoji: style code. If not set, the sparkles
emoji will be used.

If you have enabled cross-posting to Mastodon in your statuslog
settings, you can skip cross-posting to Mastodon by setting the
--skip-mastodon-post flag.`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			type Input struct {
				Emoji            string `json:"emoji"`
				Content          string `json:"content"`
				SkipMastodonPost bool   `json:"skip_mastodon_post,omitempty"`
			}
			type Result struct {
				Request struct {
					StatusCode int  `json:"status_code"`
					Success    bool `json:"success"`
				} `json:"request"`
				Response struct {
					Message     string `json:"message"`
					ID          string `json:"id"`
					Status      string `json:"status"`
					URL         string `json:"url"`
					ExternalURL string `json:"external_url"`
				} `json:"response"`
			}
			var result Result
			status := Input{
				createStatusEmoji,
				createStatusStatus,
				createStatusSkipMastodonPost,
			}
			body := callAPIWithParams(
				http.MethodPost,
				"/address/"+viper.GetString("address")+"/statuses/",
				status,
				true,
			)
			err := json.Unmarshal(body, &result)
			checkError(err)
			if !jsonFlag {
				if result.Request.Success {
					logInfo(result.Response.Message)
				} else {
					checkError(fmt.Errorf(result.Response.Message))
				}
			} else {
				fmt.Println(string(body))
			}
		},
	}
)

func init() {
	createStatusCmd.Flags().StringVarP(
		&createStatusEmoji,
		"emoji",
		"e",
		"",
		"Emoji to add to status (default sparkles)",
	)
	createStatusCmd.Flags().StringVarP(
		&createStatusStatus,
		"text",
		"t",
		"",
		"Status text",
	)
	createStatusCmd.Flags().BoolVar(
		&createStatusSkipMastodonPost,
		"skip-mastodon-post",
		false,
		"Do not cross-post to Mastodon",
	)
	createCmd.AddCommand(createStatusCmd)
}
