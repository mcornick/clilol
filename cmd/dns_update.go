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
	dnsUpdateType     string
	dnsUpdateData     string
	dnsUpdatePriority int
	dnsUpdateTTL      int
	dnsUpdateCmd      = &cobra.Command{
		Use:   "update",
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
			dns := Input{dnsUpdateType, name, dnsUpdateData, dnsUpdatePriority, dnsUpdateTTL}
			body := callAPI(
				http.MethodPatch,
				"/address/"+viper.GetString("username")+"/dns/"+objectID,
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
	dnsUpdateCmd.Flags().StringVarP(
		&objectID,
		"id",
		"i",
		"",
		"ID of DNS record to update",
	)
	dnsUpdateCmd.Flags().StringVarP(
		&dnsUpdateType,
		"type",
		"t",
		"",
		"Updated DNS type",
	)
	dnsUpdateCmd.Flags().StringVarP(
		&name,
		"name",
		"n",
		"",
		"Updated record name",
	)
	dnsUpdateCmd.Flags().StringVarP(
		&dnsUpdateData,
		"data",
		"d",
		"",
		"Updated data",
	)
	dnsUpdateCmd.Flags().IntVarP(
		&dnsUpdatePriority,
		"priority",
		"p",
		0,
		"Updated priority",
	)
	dnsUpdateCmd.Flags().IntVarP(
		&dnsUpdateTTL,
		"ttl",
		"T",
		3600,
		"Updated TTL",
	)
	dnsCmd.AddCommand(dnsUpdateCmd)
}
