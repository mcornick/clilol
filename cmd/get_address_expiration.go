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

var getAddressExpirationCmd = &cobra.Command{
	Use:   "expiration [address]",
	Short: "Get address expiration",
	Long:  "Gets the expiration of an address.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		address := args[0]
		result, err := getAddressExpiration(address)
		handleAPIError(err)
		if !result {
			fmt.Printf("%s is not expired.\n", address)
		} else {
			fmt.Printf("%s is expired.\n", address)
		}
	},
}

func init() {
	getAddressCmd.AddCommand(getAddressExpirationCmd)
}

func getAddressExpiration(address string) (bool, error) {
	client, err := omglol.NewClient(viper.GetString("email"), viper.GetString("apikey"), endpoint)
	cobra.CheckErr(err)
	expiration, err := client.GetAddressExpiration(address)
	if expiration == nil {
		return false, err
	}
	return *expiration, err
}
