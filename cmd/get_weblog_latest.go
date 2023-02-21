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

var getWeblogLatestCmd = &cobra.Command{
	Use:   "latest",
	Short: "get the latest weblog entry",
	Long:  "Gets your weblog's latest entry",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		type Result struct {
			Request struct {
				StatusCode int  `json:"status_code"`
				Success    bool `json:"success"`
			} `json:"request"`
			Response struct {
				Message string `json:"message"`
				Post    struct {
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
				} `json:"post"`
			} `json:"response"`
		}
		var result Result
		body := callAPIWithJSON(
			http.MethodGet,
			"/address/"+viper.GetString("address")+"/weblog/post/latest",
			nil,
			true,
		)
		err := json.Unmarshal(body, &result)
		cobra.CheckErr(err)
		if !silent {
			if !wantJson {
				if result.Request.Success {
					fmt.Printf(
						"%s (%s) modified on %s\n\n%s\n",
						result.Response.Post.Entry,
						fmt.Sprintf(
							"https://%s.weblog.lol%s",
							result.Response.Post.Address,
							result.Response.Post.Location,
						),
						time.Unix(result.Response.Post.Date, 0),
						result.Response.Post.Body,
					)
				} else {
					cobra.CheckErr(fmt.Errorf(result.Response.Message))
				}
			} else {
				fmt.Println(string(body))
			}
		}
	},
}

func init() {
	getWeblogCmd.AddCommand(getWeblogLatestCmd)
}
