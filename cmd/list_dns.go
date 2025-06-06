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

	"github.com/mcornick/omglol-client-go/omglol"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listDNSCmd = &cobra.Command{
	Use:   "dns",
	Short: "List your dns records",
	Long:  "Lists all of your DNS records.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		validateConfig()
		result, err := listDNS()
		handleAPIError(err)
		for _, record := range result {
			fmt.Printf(
				"%s %s %s ; ID: %d\n",
				record.Name,
				record.Type,
				record.Data,
				record.ID,
			)
		}
	},
}

func init() {
	listCmd.AddCommand(listDNSCmd)
}

func listDNS() ([]omglol.DNSRecord, error) {
	client, err := omglol.NewClient(viper.GetString("email"), viper.GetString("apikey"), endpoint)
	cobra.CheckErr(err)
	records, err := client.ListDNSRecords(viper.GetString("address"))
	return *records, err
}
