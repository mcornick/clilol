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
	dnsCreateType     string
	dnsCreateData     string
	dnsCreatePriority int
	dnsCreateTTL      int
	dnsCreateCmd      = &cobra.Command{
		Use:   "create",
		Short: "create a DNS record",
		Long: `Creates a DNS record.

Specify the type of DNS record with the --type flag,
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
			dns := Input{dnsCreateType, name, dnsCreateData, dnsCreatePriority, dnsCreateTTL}
			body := callAPIWithJSON(
				http.MethodPost,
				"/address/"+viper.GetString("address")+"/dns",
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
	dnsCreateCmd.Flags().StringVarP(
		&dnsCreateType,
		"type",
		"t",
		"",
		"Type of DNS record to create",
	)
	dnsCreateCmd.Flags().StringVarP(
		&name,
		"name",
		"n",
		"",
		"Name of the DNS record to create",
	)
	dnsCreateCmd.Flags().StringVarP(
		&dnsCreateData,
		"data",
		"d",
		"",
		"Data to store in the DNS record",
	)
	dnsCreateCmd.Flags().IntVarP(
		&dnsCreatePriority,
		"priority",
		"p",
		0,
		"Priority of the DNS record",
	)
	dnsCreateCmd.Flags().IntVarP(
		&dnsCreateTTL,
		"ttl",
		"T",
		3600,
		"Time to live of the DNS record",
	)
	dnsCmd.AddCommand(dnsCreateCmd)
}
