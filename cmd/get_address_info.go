// SPDX-License-Identifier: MPL-2.0
//
// Copyright (c) 2023 Mark Cornick
//
// This Source Code Form is subject to the terms of the Mozilla
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"fmt"

	"github.com/mcornick/omglol-client-go/omglol"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var getAddressInfoCmd = &cobra.Command{
	Use:   "info [address]",
	Short: "Get information about an address",
	Long:  "Gets information about an address.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		validateConfig()
		result, err := getAddressInfo(args[0])
		handleAPIError(err)
		fmt.Println(result.Registration.Message)
		fmt.Println(result.Expiration.Message)
		fmt.Println(result.Verification.Message)
	},
}

func init() {
	getAddressCmd.AddCommand(getAddressInfoCmd)
}

func getAddressInfo(address string) (*omglol.AddressInfo, error) {
	client, err := omglol.NewClient(viper.GetString("email"), viper.GetString("apikey"), endpoint)
	cobra.CheckErr(err)
	info, err := client.GetAddressInfo(address)
	return info, err
}
