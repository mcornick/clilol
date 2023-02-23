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

var getAddressInfoPublicCmd = &cobra.Command{
	Use:   "public",
	Short: "Get public information about an address",
	Long:  `Gets public information about an address.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		type Result struct {
			Request struct {
				StatusCode int  `json:"status_code"`
				Success    bool `json:"success"`
			} `json:"request"`
			Response struct {
				Address      string `json:"address"`
				Message      string `json:"message"`
				Registration struct {
					Message       string    `json:"message"`
					UnixEpochTime int64     `json:"unix_epoch_time"`
					ISO8601Time   time.Time `json:"iso_8601_time"`
					RFC2822Time   string    `json:"rfc_2822_time"`
					RelativeTime  string    `json:"relative_time"`
				} `json:"registration"`
				Expiration struct {
					Message string `json:"message"`
					Expired bool   `json:"expired"`
				} `json:"expiration"`
				Verification struct {
					Message  string `json:"message"`
					Verified bool   `json:"verified"`
				} `json:"verification"`
			} `json:"response"`
		}
		var result Result
		if addressFlag == "" {
			addressFlag = viper.GetString("address")
		}
		body := callAPIWithParams(
			http.MethodGet,
			"/address/"+addressFlag+"/info",
			nil,
			false,
		)
		err := json.Unmarshal(body, &result)
		checkError(err)
		if !silentFlag {
			if !jsonFlag {
				if result.Request.Success {
					fmt.Println(result.Response.Registration.Message)
					fmt.Println(result.Response.Expiration.Message)
					fmt.Println(result.Response.Verification.Message)
				} else {
					checkError(fmt.Errorf(result.Response.Message))
				}
			} else {
				fmt.Println(string(body))
			}
		}
	},
}

func init() {
	getAddressInfoPublicCmd.Flags().StringVarP(
		&addressFlag,
		"address",
		"a",
		"",
		"address whose availability to get",
	)
	getAddressInfoCmd.AddCommand(getAddressInfoPublicCmd)
}
