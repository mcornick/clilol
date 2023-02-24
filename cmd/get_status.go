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

var getStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get status",
	Long: `Gets a single status for a single user from status.lol.

Specify the status ID with the --id flag.

The address can be specified with the --address flag. If not set,
it defaults to your own address.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		type Result struct {
			Request  responseRequest `json:"request"`
			Response struct {
				Message string `json:"message"`
				Status  struct {
					Id      string `json:"id"`
					Address string `json:"address"`
					Created string `json:"created"`
					Emoji   string `json:"emoji"`
					Content string `json:"content"`
				} `json:"status"`
			} `json:"response"`
		}
		var result Result
		if addressFlag == "" {
			addressFlag = viper.GetString("address")
		}
		body := callAPIWithParams(
			http.MethodGet,
			"/address/"+addressFlag+"/statuses/"+idFlag,
			nil,
			false,
		)
		err := json.Unmarshal(body, &result)
		cobra.CheckErr(err)
		if result.Request.Success {
			fmt.Printf(
				"\nhttps://status.lol/%s/%s\n",
				result.Response.Status.Address,
				result.Response.Status.Id,
			)
			timestamp, err := strconv.Atoi(result.Response.Status.Created)
			cobra.CheckErr(err)
			fmt.Printf("  %s\n", time.Unix(int64(timestamp), 0))
			fmt.Printf(
				"  %s %s\n",
				result.Response.Status.Emoji,
				result.Response.Status.Content,
			)
		} else {
			fmt.Println(string(body))
		}
	},
}

func init() {
	getStatusCmd.Flags().StringVarP(
		&addressFlag,
		"address",
		"a",
		"",
		"address whose status to get",
	)
	getStatusCmd.Flags().StringVarP(
		&idFlag,
		"id",
		"i",
		"",
		"ID of the status to get",
	)
	err := getStatusCmd.MarkFlagRequired("id")
	cobra.CheckErr(err)
	getCmd.AddCommand(getStatusCmd)
}
