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

var getAccountInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info about your account",
	Long:  "Gets information about your account",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		type Result struct {
			Request  responseRequest `json:"request"`
			Response struct {
				Message string `json:"message"`
				Email   string `json:"email"`
				Name    string `json:"name"`
				Created struct {
					UnixEpochTime int64     `json:"unix_epoch_time"`
					ISO8601Time   time.Time `json:"iso_8601_time"`
					RFC2822Time   string    `json:"rfc_2822_time"`
					RelativeTime  string    `json:"relative_time"`
				} `json:"created"`
				Settings struct {
					Communication string `json:"communication"`
				} `json:"settings"`
			} `json:"response"`
		}
		var result Result
		body := callAPIWithParams(
			http.MethodGet,
			"/account/"+viper.GetString("email")+"/info",
			nil,
			true,
		)
		err := json.Unmarshal(body, &result)
		cobra.CheckErr(err)
		if result.Request.Success {
			fmt.Println(result.Response.Message)
			fmt.Printf("%s (%s)\n", result.Response.Name, result.Response.Email)
			fmt.Printf("Created %s\n", result.Response.Created.RelativeTime)
			fmt.Printf("Communication: %s\n", result.Response.Settings.Communication)
		} else {
			cobra.CheckErr(fmt.Errorf(result.Response.Message))
		}
	},
}

func init() {
	getAccountCmd.AddCommand(getAccountInfoCmd)
}
