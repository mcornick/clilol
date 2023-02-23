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

var getServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Get service stats",
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
		body := callAPIWithParams(http.MethodGet, "/service/info", nil, false)
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

func init() {
	getCmd.AddCommand(getServiceCmd)
}
