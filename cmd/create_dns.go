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

	"github.com/ejstreet/omglol-client-go/omglol"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	createDNSPriority int64
	createDNSTTL      int64
	createDNSCmd      = &cobra.Command{
		Use:   "dns [name] [type] [data]",
		Short: "Create a DNS record",
		Long:  "Creates a DNS record.",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			validateConfig()
			name := args[0]
			recordType := args[1]
			data := args[2]
			err := createDNS(name, recordType, data, createDNSPriority, createDNSTTL)
			handleAPIError(err)
			fmt.Printf("DNS record %s created.\n", name)
		},
	}
)

func init() {
	createDNSCmd.Flags().Int64VarP(
		&createDNSPriority,
		"priority",
		"p",
		0,
		"priority of the DNS record",
	)
	createDNSCmd.Flags().Int64VarP(
		&createDNSTTL,
		"ttl",
		"T",
		3600,
		"time to live of the DNS record",
	)
	createCmd.AddCommand(createDNSCmd)
}

func createDNS(name string, recordType string, data string, priority int64, ttl int64) error {
	client, err := omglol.NewClient(viper.GetString("email"), viper.GetString("apikey"), endpoint)
	cobra.CheckErr(err)
	dnsEntry := omglol.NewDNSEntry(recordType, name, data, ttl, priority)
	_, err = client.CreateDNSRecord(viper.GetString("address"), *dnsEntry)
	return err
}
