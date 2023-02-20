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
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	postEmoji       string
	postExternalURL string
	postCmd         = &cobra.Command{
		Use:   "post [status text]",
		Short: "post a status",
		Long: `Posts a status to status.lol.

Quote the status text if it contains spaces.

You can specify an emoji with the --emoji flag. This must be an
actual emoji, not a :emoji: style code. If not set, the sparkles
emoji will be used.

You can specify an external URL with the --external-url flag. This
will be shown as a "Respond" link on the statuslog. If not set, no
external URL will be used.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			type Input struct {
				Emoji       string `json:"emoji"`
				Content     string `json:"content"`
				ExternalURL string `json:"external_url,omitempty"`
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
			status := Input{postEmoji, strings.Join(args, " "), postExternalURL}
			body := callAPI(
				http.MethodPost,
				"/address/"+viper.GetString("username")+"/statuses/",
				status,
				true,
			)
			err := json.Unmarshal(body, &result)
			cobra.CheckErr(err)
			if !silent {
				if !wantJson {
					if result.Request.Success {
						fmt.Println(result.Response.Message)
					} else {
						cobra.CheckErr(fmt.Errorf(result.Response.Message))
					}
				} else {
					fmt.Println(string(body))
				}
			}

		},
	}
)

func init() {
	postCmd.Flags().StringVarP(
		&postEmoji,
		"emoji",
		"e",
		"",
		"Emoji to add to status (default sparkles)",
	)
	postCmd.Flags().StringVarP(
		&postExternalURL,
		"external-url",
		"u",
		"",
		"External URL to add to status",
	)
	statusCmd.AddCommand(postCmd)
}
