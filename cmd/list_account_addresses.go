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

var listAccountAddressesCmd = &cobra.Command{
	Use:     "addresses",
	Aliases: []string{"address"},
	Short:   "List addresses on your account",
	Long:    "Lists addresses on your account.",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		validateConfig()
		addresses, err := listAccountAddresses()
		handleAPIError(err)
		for _, address := range addresses {
			fmt.Println(address.Address)
		}
	},
}

func init() {
	listAccountCmd.AddCommand(listAccountAddressesCmd)
}

func listAccountAddresses() ([]omglol.Address, error) {
	client, err := omglol.NewClient(viper.GetString("email"), viper.GetString("apikey"), endpoint)
	cobra.CheckErr(err)
	addresses, err := client.GetAccountAddresses()
	return *addresses, err
}
