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

var listDNSCmd = &cobra.Command{
	Use:   "dns",
	Short: "List your dns records",
	Long:  "Lists all of your DNS records.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		type output struct {
			Request  resultRequest `json:"request"`
			Response struct {
				Message string `json:"message"`
				DNS     []struct {
					ID        int       `json:"id"`
					Type      string    `json:"type"`
					Name      string    `json:"name"`
					Data      string    `json:"data"`
					Priority  int       `json:"priority"`
					TTL       int       `json:"ttl"`
					CreatedAt time.Time `json:"created_at"`
					UpdatedAt time.Time `json:"updated_at"`
				} `json:"dns"`
			} `json:"response"`
		}
		var result output
		body := callAPIWithParams(
			http.MethodGet,
			"/address/"+viper.GetString("address")+"/dns",
			nil,
			true,
		)
		err := json.Unmarshal(body, &result)
		cobra.CheckErr(err)
		if result.Request.Success {
			for _, record := range result.Response.DNS {
				fmt.Printf(
					"%s %s %s ; ID: %d\n",
					record.Name,
					record.Type,
					record.Data,
					record.ID,
				)
			}
		} else {
			cobra.CheckErr(fmt.Errorf(result.Response.Message))
		}
	},
}

func init() {
	listCmd.AddCommand(listDNSCmd)
}
