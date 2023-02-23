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

var (
	listStatuslogAll bool
	listStatuslogCmd = &cobra.Command{
		Use:   "statuslog",
		Short: "List the statuslog",
		Long: `Lists status(es) for all status.lol users.

By default, only the most recent status for each user is returned.
To see all statuses ever posted, use the --all flag.

See the status commands to get statuses for a single user.`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			type Result struct {
				Request  responseRequest `json:"request"`
				Response struct {
					Message  string `json:"message"`
					Statuses []struct {
						Id           string `json:"id"`
						Address      string `json:"address"`
						Created      string `json:"created"`
						RelativeTime string `json:"relative_time"`
						Emoji        string `json:"emoji"`
						Content      string `json:"content"`
					} `json:"statuses"`
				} `json:"response"`
			}
			var url string
			if listStatuslogAll {
				url = "/statuslog/"
			} else {
				url = "/statuslog/latest/"
			}
			var result Result
			body := callAPIWithParams(http.MethodGet, url, nil, false)
			err := json.Unmarshal(body, &result)
			checkError(err)
			if result.Request.Success {
				for _, status := range result.Response.Statuses {
					fmt.Printf("@%s, %s\n", status.Address, status.RelativeTime)
					fmt.Printf("  %s %s\n", status.Emoji, status.Content)
				}
			} else {
				checkError(fmt.Errorf(result.Response.Message))
			}
		},
	}
)

func init() {
	listStatuslogCmd.Flags().BoolVarP(
		&listStatuslogAll,
		"all",
		"A",
		false,
		"get the entire statuslog (default is latest statuses only)",
	)
	listCmd.AddCommand(listStatuslogCmd)
}
