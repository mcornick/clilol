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
	"github.com/spf13/viper"
)

var getWeblogCmd = &cobra.Command{
	Use:   "weblog",
	Short: "Get a weblog entry",
	Long: `Gets one of your weblog entries by ID.

Specify the ID with the --id flag.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		type Result struct {
			Request  responseRequest `json:"request"`
			Response struct {
				Message string `json:"message"`
				Entry   struct {
					Address  string `json:"address"`
					Location string `json:"location"`
					Title    string `json:"title"`
					Date     int64  `json:"date"`
					Type     string `json:"type"`
					Status   string `json:"status"`
					Source   string `json:"source"`
					Body     string `json:"body"`
					Output   string `json:"output"`
					Metadata string `json:"metadata"`
					Entry    string `json:"entry"`
					ID       string `json:"id"`
				} `json:"entry"`
			} `json:"response"`
		}
		var result Result
		body := callAPIWithParams(
			http.MethodGet,
			"/address/"+viper.GetString("address")+"/weblog/entry/"+idFlag,
			nil,
			true,
		)
		err := json.Unmarshal(body, &result)
		checkError(err)
		if result.Request.Success {
			fmt.Printf(
				"%s (%s) modified on %s\n\n%s\n",
				result.Response.Entry.Entry,
				fmt.Sprintf(
					"https://%s.weblog.lol%s",
					result.Response.Entry.Address,
					result.Response.Entry.Location,
				),
				time.Unix(result.Response.Entry.Date, 0),
				result.Response.Entry.Body,
			)
		} else {
			checkError(fmt.Errorf(result.Response.Message))
		}
	},
}

func init() {
	getWeblogCmd.Flags().StringVarP(
		&idFlag,
		"id",
		"i",
		"",
		"ID of the entry to get",
	)
	getWeblogCmd.MarkFlagRequired("id")
	getCmd.AddCommand(getWeblogCmd)
}
