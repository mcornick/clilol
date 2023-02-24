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

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	createDNSType     string
	createDNSData     string
	createDNSPriority int
	createDNSTTL      int
	createDNSCmd      = &cobra.Command{
		Use:   "dns",
		Short: "Create a DNS record",
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
				Request  responseRequest `json:"request"`
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
			dns := Input{createDNSType, nameFlag, createDNSData, createDNSPriority, createDNSTTL}
			body := callAPIWithParams(
				http.MethodPost,
				"/address/"+viper.GetString("address")+"/dns",
				dns,
				true,
			)
			err := json.Unmarshal(body, &result)
			checkError(err)
			if result.Request.Success {
				log.Info(result.Response.Message)
			} else {
				checkError(fmt.Errorf(result.Response.Message))
			}
		},
	}
)

func init() {
	createDNSCmd.Flags().StringVarP(
		&createDNSType,
		"type",
		"t",
		"",
		"type of DNS record to create",
	)
	createDNSCmd.Flags().StringVarP(
		&nameFlag,
		"name",
		"n",
		"",
		"name of the DNS record to create",
	)
	createDNSCmd.Flags().StringVarP(
		&createDNSData,
		"data",
		"d",
		"",
		"data to store in the DNS record",
	)
	createDNSCmd.Flags().IntVarP(
		&createDNSPriority,
		"priority",
		"p",
		0,
		"priority of the DNS record",
	)
	createDNSCmd.Flags().IntVarP(
		&createDNSTTL,
		"ttl",
		"T",
		3600,
		"time to live of the DNS record",
	)
	createDNSCmd.MarkFlagRequired("type")
	createDNSCmd.MarkFlagRequired("name")
	createDNSCmd.MarkFlagRequired("data")
	createCmd.AddCommand(createDNSCmd)
}
