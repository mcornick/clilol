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

var accountListAddressesCmd = &cobra.Command{
	Use:   "addresses",
	Short: "list your addresses",
	Long:  `Lists the addresses on your account.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		type Result struct {
			Request struct {
				StatusCode int  `json:"status_code"`
				Success    bool `json:"success"`
			} `json:"request"`
			Response []struct {
				Address      string `json:"address"`
				Message      string `json:"message"`
				Registration struct {
					Message       string    `json:"message"`
					UnixEpochTime int       `json:"unix_epoch_time"`
					ISO8601Time   time.Time `json:"iso_8601_time"`
					RFC2822Time   string    `json:"rfc_2822_time"`
					RelativeTime  string    `json:"relative_time"`
				} `json:"registration"`
				Expiration struct {
					Expired       bool      `json:"expired"`
					WillExpire    bool      `json:"will_expire"`
					UnixEpochTime int       `json:"unix_epoch_time"`
					ISO8601Time   time.Time `json:"iso_8601_time"`
					RFC2822Time   string    `json:"rfc_2822_time"`
					RelativeTime  string    `json:"relative_time"`
				} `json:"expiration"`
				Preferences struct {
					IncludeInDirectory string `json:"include_in_directory"`
					ShowOnDashboard    string `json:"show_on_dashboard"`
					Statuslog          struct {
						MastodonPosting bool `json:"mastodon_posting"`
					} `json:"statuslog"`
				} `json:"preferences"`
			} `json:"response"`
		}
		var result Result
		body := callAPI(
			http.MethodGet,
			"/account/"+viper.GetString("email")+"/addresses",
			nil,
			true,
		)
		err := json.Unmarshal(body, &result)
		cobra.CheckErr(err)
		if !silent {
			if !wantJson {
				if result.Request.Success {
					for _, address := range result.Response {
						fmt.Println(address.Address)
						fmt.Println(address.Message)
						fmt.Printf("Registered %s\n", address.Registration.RelativeTime)
					}
				} else {
					cobra.CheckErr(fmt.Errorf("%d", result.Request.StatusCode))
				}
			} else {
				fmt.Println(string(body))
			}
		}
	},
}

func init() {
	accountListCmd.AddCommand(accountListAddressesCmd)
}