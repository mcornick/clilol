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
	"github.com/spf13/viper"
)

var (
	listStatusLimit int
	listStatusCmd   = &cobra.Command{
		Use:   "status",
		Short: "List statuses",
		Long: `Lists statuses for a single user from status.lol.

The address can be specified with the --address flag. If not set,
it defaults to your own address.

The number of statuses returned can be specified with the --limit
flag. If not set, it will return all statuses for the user.

See the statuslog commands to get statuses for all users.`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			type Result struct {
				Request  resultRequest `json:"request"`
				Response struct {
					Message  string `json:"message"`
					Statuses []struct {
						Id      string `json:"id"`
						Address string `json:"address"`
						Created string `json:"created"`
						Emoji   string `json:"emoji"`
						Content string `json:"content"`
					} `json:"statuses"`
				} `json:"response"`
			}
			var result Result
			if addressFlag == "" {
				addressFlag = viper.GetString("address")
			}
			body := callAPIWithParams(
				http.MethodGet,
				"/address/"+addressFlag+"/statuses/",
				nil,
				false,
			)
			err := json.Unmarshal(body, &result)
			cobra.CheckErr(err)
			if listStatusLimit > 0 {
				result.Response.Statuses = result.Response.Statuses[:listStatusLimit]
			}
			if result.Request.Success {
				for _, status := range result.Response.Statuses {
					fmt.Printf("\nhttps://status.lol/%s/%s\n", status.Address, status.Id)
					timestamp, err := strconv.Atoi(status.Created)
					cobra.CheckErr(err)
					fmt.Printf("  %s\n", time.Unix(int64(timestamp), 0))
					fmt.Printf("  %s %s\n", status.Emoji, status.Content)
				}
			} else {
				cobra.CheckErr(fmt.Errorf(result.Response.Message))
			}
		},
	}
)

func init() {
	listStatusCmd.Flags().StringVarP(
		&addressFlag,
		"address",
		"a",
		"",
		"address whose status(es) to get",
	)
	listStatusCmd.Flags().IntVarP(
		&listStatusLimit,
		"limit",
		"l",
		0,
		"how many status(es) to get (default all)",
	)
	listCmd.AddCommand(listStatusCmd)
}
