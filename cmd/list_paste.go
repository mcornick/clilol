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

var listPasteCmd = &cobra.Command{
	Use:   "paste",
	Short: "List pastes",
	Long: `Lists pastes for a user.

The address can be specified with the --address flag. If not set,
it defaults to your own address.

Unlisted pastes are only included when the --address flag is set to
your own address.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		type Result struct {
			Request struct {
				StatusCode int  `json:"status_code"`
				Success    bool `json:"success"`
			} `json:"request"`
			Response struct {
				Message  string `json:"message"`
				Pastebin []struct {
					Title      string `json:"title"`
					Content    string `json:"content"`
					ModifiedOn int64  `json:"modified_on"`
					Listed     int    `json:"listed,omitempty"`
				} `json:"pastebin"`
			} `json:"response"`
		}
		var result Result
		if addressFlag == "" {
			addressFlag = viper.GetString("address")
		}
		body := callAPIWithParams(
			http.MethodGet,
			"/address/"+addressFlag+"/pastebin",
			nil,
			addressFlag == viper.GetString("address"),
		)
		err := json.Unmarshal(body, &result)
		checkError(err)
		if !jsonFlag {
			if result.Request.Success {
				for _, paste := range result.Response.Pastebin {
					fmt.Printf(
						"%s modified on %s\n",
						paste.Title,
						time.Unix(paste.ModifiedOn, 0),
					)
				}
			} else {
				checkError(fmt.Errorf(result.Response.Message))
			}
		} else {
			fmt.Println(string(body))
		}
	},
}

func init() {
	listPasteCmd.Flags().StringVarP(
		&addressFlag,
		"address",
		"a",
		"",
		"address whose pastes to list",
	)
	listCmd.AddCommand(listPasteCmd)
}
