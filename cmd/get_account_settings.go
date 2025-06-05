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

var getAccountSettingsCmd = &cobra.Command{
	Use:   "settings",
	Short: "Get your account settings",
	Long:  "Gets the settings on your account.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// var result getAccountSettingsOutput
		result, err := getAccountSettings()
		handleAPIError(err)
		fmt.Printf("Owner: %s\n", result.Owner)
		if result.Communication != nil {
			fmt.Printf("Communication: %s\n", *result.Communication)
		}
		if result.DateFormat != nil {
			fmt.Printf("Date Format: %s\n", *result.DateFormat)
		}
		if result.WebEditor != nil {
			fmt.Printf("Web Editor: %s\n", *result.WebEditor)
		}
	},
}

func init() {
	getAccountCmd.AddCommand(getAccountSettingsCmd)
}

func getAccountSettings() (omglol.AccountSettings, error) {
	client, err := omglol.NewClient(viper.GetString("email"), viper.GetString("apikey"), endpoint)
	cobra.CheckErr(err)
	settings, err := client.GetAccountSettings()
	return *settings, err
}
