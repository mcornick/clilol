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
	"github.com/spf13/viper"
)

var (
	updateDNSType     string
	updateDNSData     string
	updateDNSPriority int
	updateDNSTTL      int
	updateDNSCmd      = &cobra.Command{
		Use:   "dns",
		Short: "update a DNS record",
		Long: `Updates a DNS record.

Specify the ID of the DNS record with the --id flag,
the type of DNS record with the --type flag,
the name of the record with the --name flag,
and the data with the --data flag.`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			type Input struct {
				Type     string `json:"type"`
				Name     string `json:"name"`
				Data     string `json:"data"`
				Priority int    `json:"priority"`
				TTL      int    `json:"ttl"`
			}
			type Result struct {
				Request struct {
					StatusCode int  `json:"status_code"`
					Success    bool `json:"success"`
				} `json:"request"`
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
							ID        int    `json:"id"`
							Name      string `json:"name"`
							Content   string `json:"content"`
							TTL       int    `json:"ttl"`
							Priority  int    `json:"priority"`
							Type      string `json:"type"`
							CreatedAt string `json:"created_at"`
							UpdatedAt string `json:"updated_at"`
						} `json:"data"`
					} `json:"response_received"`
				} `json:"response"`
			}
			var result Result
			dns := Input{updateDNSType, nameFlag, updateDNSData, updateDNSPriority, updateDNSTTL}
			body := callAPIWithJSON(
				http.MethodPatch,
				"/address/"+viper.GetString("address")+"/dns/"+idFlag,
				dns,
				true,
			)
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
)

func init() {
	updateDNSCmd.Flags().StringVarP(
		&idFlag,
		"id",
		"i",
		"",
		"ID of DNS record to update",
	)
	updateDNSCmd.Flags().StringVarP(
		&updateDNSType,
		"type",
		"t",
		"",
		"Updated DNS type",
	)
	updateDNSCmd.Flags().StringVarP(
		&nameFlag,
		"name",
		"n",
		"",
		"Updated record name",
	)
	updateDNSCmd.Flags().StringVarP(
		&updateDNSData,
		"data",
		"d",
		"",
		"Updated data",
	)
	updateDNSCmd.Flags().IntVarP(
		&updateDNSPriority,
		"priority",
		"p",
		0,
		"Updated priority",
	)
	updateDNSCmd.Flags().IntVarP(
		&updateDNSTTL,
		"ttl",
		"T",
		3600,
		"Updated TTL",
	)
	updateCmd.AddCommand(updateDNSCmd)
}
