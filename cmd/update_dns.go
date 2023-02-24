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

var (
	updateDNSPriority int
	updateDNSTTL      int
	updateDNSCmd      = &cobra.Command{
		Use:   "dns [id] [name] [type] [data]",
		Short: "Update a DNS record",
		Long:  "Updates a DNS record.",
		Args:  cobra.ExactArgs(4),
		Run: func(cmd *cobra.Command, args []string) {
			type Input struct {
				Type     string `json:"type"`
				Name     string `json:"name"`
				Data     string `json:"data"`
				Priority int    `json:"priority"`
				TTL      int    `json:"ttl"`
			}
			type Result struct {
				Request  resultRequest `json:"request"`
				Response struct {
					Message  string `json:"message"`
					DataSent struct {
						Type     string `json:"type"`
						Priority int    `json:"priority"`
						TTL      int    `json:"ttl"`
						Name     string `json:"name"`
						Content  string `json:"content"`
					} `json:"data_sent"`
					ResponseReceived struct {
						Data struct {
							ID        int       `json:"id"`
							Name      string    `json:"name"`
							Content   string    `json:"content"`
							TTL       int       `json:"ttl"`
							Priority  int       `json:"priority"`
							Type      string    `json:"type"`
							CreatedAt time.Time `json:"created_at"`
							UpdatedAt time.Time `json:"updated_at"`
						} `json:"data"`
					} `json:"response_received"`
				} `json:"response"`
			}
			var result Result
			dns := Input{args[2], args[1], args[3], updateDNSPriority, updateDNSTTL}
			body := callAPIWithParams(
				http.MethodPatch,
				"/address/"+viper.GetString("address")+"/dns/"+args[0],
				dns,
				true,
			)
			err := json.Unmarshal(body, &result)
			cobra.CheckErr(err)
			if result.Request.Success {
				fmt.Println(result.Response.Message)
			} else {
				cobra.CheckErr(fmt.Errorf(result.Response.Message))
			}
		},
	}
)

func init() {
	updateDNSCmd.Flags().IntVarP(
		&updateDNSPriority,
		"priority",
		"p",
		0,
		"updated priority",
	)
	updateDNSCmd.Flags().IntVarP(
		&updateDNSTTL,
		"ttl",
		"T",
		3600,
		"updated TTL",
	)
	updateCmd.AddCommand(updateDNSCmd)
}
