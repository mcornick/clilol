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

var serviceInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "get service stats",
	Long:  "Gets statistics for omg.lol services.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		type Result struct {
			Request struct {
				StatusCode int  `json:"status_code"`
				Success    bool `json:"success"`
			} `json:"request"`
			Response struct {
				Message   string `json:"message"`
				Members   int    `json:"members"`
				Addresses int    `json:"addresses"`
				Profiles  int    `json:"profiles"`
			} `json:"response"`
		}
		var result Result
		body := callAPI(http.MethodGet, "/service/info", nil, false)
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

func init() {
	serviceCmd.AddCommand(serviceInfoCmd)
}
