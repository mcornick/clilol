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
	"strconv"

	"github.com/ejstreet/omglol-client-go/omglol"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	updateDNSPriority int64
	updateDNSTTL      int64
	updateDNSCmd      = &cobra.Command{
		Use:   "dns [id] [name] [type] [data]",
		Short: "Update a DNS record",
		Long:  "Updates a DNS record.",
		Args:  cobra.ExactArgs(4),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.ParseInt(args[0], 10, 64)
			cobra.CheckErr(err)
			name := args[1]
			recordType := args[2]
			data := args[3]
			err = updateDNS(id, name, recordType, data, updateDNSPriority, updateDNSTTL)
			handleAPIError(err)
			fmt.Printf("DNS record %s updated.\n", name)
		},
	}
)

func init() {
	updateDNSCmd.Flags().Int64VarP(
		&updateDNSPriority,
		"priority",
		"p",
		0,
		"updated priority",
	)
	updateDNSCmd.Flags().Int64VarP(
		&updateDNSTTL,
		"ttl",
		"T",
		3600,
		"updated TTL",
	)
	updateCmd.AddCommand(updateDNSCmd)
}

func updateDNS(id int64, name string, recordType string, data string, priority int64, ttl int64) error {
	client, err := omglol.NewClient(viper.GetString("email"), viper.GetString("apikey"), endpoint)
	cobra.CheckErr(err)
	dnsEntry := omglol.NewDNSEntry(recordType, name, data, ttl, priority)
	_, err = client.UpdateDNSRecord(viper.GetString("address"), *dnsEntry, id)
	return err
}
