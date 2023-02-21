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
	updateStatusEmoji  string
	updateStatusStatus string
	updateStatusCmd    = &cobra.Command{
		Use:   "status",
		Short: "update a status",
		Long: `Updates a status on status.lol.
Specify the ID of the status to update with the --id flag. The
status can be found as the last element of the status URL.

Specify the new status text with the --text flag.
Quote the text if it contains spaces.

You can specify an emoji with the --emoji flag. This must be an
actual emoji, not a :emoji: style code. If not set, the sparkles
emoji will be used. Note that the omg.lol API does not preserve
the existing emoji if you don't specify one, so if you don't want
to change it, you'll still need to specify it again.`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			type Input struct {
				Id      string `json:"id"`
				Emoji   string `json:"emoji"`
				Content string `json:"content"`
			}
			type Result struct {
				Request struct {
					StatusCode int  `json:"status_code"`
					Success    bool `json:"success"`
				} `json:"request"`
				Response struct {
					Message string `json:"message"`
					ID      string `json:"id"`
					URL     string `json:"url"`
				} `json:"response"`
			}
			var result Result
			status := Input{idFlag, updateStatusEmoji, updateStatusStatus}
			body := callAPIWithJSON(
				http.MethodPatch,
				"/address/"+viper.GetString("address")+"/statuses/",
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
	updateStatusCmd.Flags().StringVarP(
		&idFlag,
		"id",
		"i",
		"",
		"ID of the status to update",
	)
	updateStatusCmd.Flags().StringVarP(
		&updateStatusEmoji,
		"emoji",
		"e",
		"",
		"Emoji to add to status (default sparkles)",
	)
	updateStatusCmd.Flags().StringVarP(
		&updateStatusStatus,
		"text",
		"t",
		"",
		"New status text",
	)
	updateCmd.AddCommand(updateStatusCmd)
}