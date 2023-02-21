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
	"time"

	"github.com/spf13/cobra"
)

var (
	nowListCmd = &cobra.Command{
		Use:   "list",
		Short: "list Now pages",
		Long:  "Lists pages in the Now garden.",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			type Result struct {
				Request struct {
					StatusCode int  `json:"status_code"`
					Success    bool `json:"success"`
				} `json:"request"`
				Response struct {
					Message string `json:"message"`
					Garden  []struct {
						Address string `json:"address"`
						URL     string `json:"url"`
						Updated struct {
							UnixEpochTime int64     `json:"unix_epoch_time"`
							ISO8601Time   time.Time `json:"iso_8601_time"`
							RFC2822Time   string    `json:"rfc_2822_time"`
							RelativeTime  string    `json:"relative_time"`
						} `json:"updated"`
					} `json:"garden"`
				} `json:"response"`
			}
			var result Result
			body := callAPI(http.MethodGet, "/now/garden", nil, false)
			err := json.Unmarshal(body, &result)
			cobra.CheckErr(err)
			if !silent {
				if !wantJson {
					if result.Request.Success {
						fmt.Println(result.Response.Message)
						for _, page := range result.Response.Garden {
							fmt.Printf("\n%s (%s)\n", page.URL, page.Address)
							fmt.Printf("last updated %s\n", page.Updated.RelativeTime)
						}
					} else {
						fmt.Println(result.Response.Message)
					}
				} else {
					cobra.CheckErr(fmt.Errorf(result.Response.Message))
				}
			} else {
				fmt.Println(string(body))
			}
		},
	}
)

func init() {
	nowCmd.AddCommand(nowListCmd)
}
