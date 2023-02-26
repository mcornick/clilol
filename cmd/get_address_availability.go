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

var getAddressAvailabilityCmd = &cobra.Command{
	Use:   "availability [address]",
	Short: "Get address availability",
	Long:  "Gets the availability of an address.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		address := args[0]
		result, err := getAddressAvailability(address)
		handleAPIError(err)
		fmt.Printf("%s is %s\n", address, result.Availability)
	},
}

func init() {
	getAddressCmd.AddCommand(getAddressAvailabilityCmd)
}

func getAddressAvailability(address string) (*omglol.AddressAvailability, error) {
	client, err := omglol.NewClient(viper.GetString("email"), viper.GetString("apikey"), endpoint)
	cobra.CheckErr(err)
	availability, err := client.GetAddressAvailability(address)
	return availability, err
}
