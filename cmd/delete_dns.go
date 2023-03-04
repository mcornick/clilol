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

var deleteDNSCmd = &cobra.Command{
	Use:   "dns [id]",
	Short: "Delete a DNS record",
	Long: `Deletes a DNS record.

Note that you won't be asked to confirm deletion.
Be sure you know what you're doing.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		validateConfig()
		id, err := strconv.ParseInt(args[0], 10, 64)
		cobra.CheckErr(err)
		err = deleteDNS(id)
		handleAPIError(err)
		fmt.Printf("DNS record %d deleted\n", id)
	},
}

func init() {
	deleteCmd.AddCommand(deleteDNSCmd)
}

func deleteDNS(id int64) error {
	client, err := omglol.NewClient(viper.GetString("email"), viper.GetString("apikey"), endpoint)
	cobra.CheckErr(err)
	err = client.DeleteDNSRecord(viper.GetString("address"), id)
	return err
}
