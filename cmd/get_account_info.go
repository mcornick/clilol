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

var getAccountInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info about your account",
	Long:  "Gets information about your account",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		validateConfig()
		result, err := getAccountInfo()
		handleAPIError(err)
		fmt.Printf("%s (%s)\n", result.Name, result.Email)
		fmt.Printf("Created %s\n", result.Created.RelativeTime)
		fmt.Printf("Communication: %s\n", *result.Settings.Communication)
	},
}

func init() {
	getAccountCmd.AddCommand(getAccountInfoCmd)
}

func getAccountInfo() (omglol.Account, error) {
	client, err := omglol.NewClient(viper.GetString("email"), viper.GetString("apikey"), endpoint)
	cobra.CheckErr(err)
	account, err := client.GetAccountInfo()
	return *account, err
}
