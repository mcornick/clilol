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

var dnsListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all dns records",
	Long:  `Lists all of your DNS records.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		type Result struct {
			Request struct {
				StatusCode int  `json:"status_code"`
				Success    bool `json:"success"`
			} `json:"request"`
			Response struct {
				Message string `json:"message"`
				DNS     []struct {
					ID         string `json:"id"`
					Type       string `json:"type"`
					Name       string `json:"name"`
					Data       string `json:"data"`
					Priority   string `json:"priority"`
					TTL        string `json:"ttl"`
					CreatedAt  string `json:"created_at"`
					Updated_At string `json:"updated_at"`
				} `json:"dns"`
			} `json:"response"`
		}
		var result Result
		body := callAPI(http.MethodGet, "/address/"+username+"/dns", nil, true)
		err := json.Unmarshal(body, &result)
		cobra.CheckErr(err)
		if !silent {
			if !wantJson {
				if result.Request.Success {
					for _, record := range result.Response.DNS {
						fmt.Printf(
							"%s %s %s\n",
							record.Name,
							record.Type,
							record.Data,
						)
					}
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
	dnsCmd.AddCommand(dnsListCmd)
}
