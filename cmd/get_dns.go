// SPDX-License-Identifier: MPL-2.0
//
// Copyright (c) 2023 Mark Cornick
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"fmt"
	"strings"

	"github.com/mcornick/omglol-client-go/omglol"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	getDNSPriority int64
	getDNSTTL      int64
	getDNSCmd      = &cobra.Command{
		Use:   "dns [name] [type] [data]",
		Short: "Get a DNS record",
		Long:  "Gets a DNS record by attributes.",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) { // result, err := getDNS(args[0], args[1], args[2], getDNSPriority, getDNSTTL)
			name := args[0]
			recordType := args[1]
			data := args[2]
			if !strings.Contains(name, ".") {
				name = name + "." + viper.GetString("address")
			}
			record, err := getDNS(name, recordType, data, getDNSPriority, getDNSTTL)
			cobra.CheckErr(err)
			fmt.Printf(
				"%s %s %s ; ID: %d\n",
				record.Name,
				record.Type,
				record.Data,
				record.ID,
			)
		},
	}
)

func init() {
	getDNSCmd.Flags().Int64VarP(
		&getDNSPriority,
		"priority",
		"p",
		0,
		"priority of the DNS record",
	)
	getDNSCmd.Flags().Int64VarP(
		&getDNSTTL,
		"ttl",
		"T",
		3600,
		"time to live of the DNS record",
	)
	getCmd.AddCommand(getDNSCmd)
}

func getDNS(name string, recordType string, data string, priority int64, ttl int64) (omglol.DNSRecord, error) {
	client, err := omglol.NewClient(viper.GetString("email"), viper.GetString("apikey"), endpoint)
	cobra.CheckErr(err)
	criteria := map[string]any{
		"Name":     name,
		"Type":     strings.ToUpper(recordType),
		"Data":     data,
		"Priority": priority,
		"TTL":      ttl,
	}
	record, err := client.FilterDNSRecord(viper.GetString("address"), criteria)
	if err != nil {
		return omglol.DNSRecord{}, err
	}
	return *record, err
}
