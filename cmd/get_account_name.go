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

var getAccountNameCmd = &cobra.Command{
	Use:   "name",
	Short: "Get your account name",
	Long:  "Gets the name on your account.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		validateConfig()
		name, err := getAccountName()
		handleAPIError(err)
		fmt.Println(name)
	},
}

func init() {
	getAccountCmd.AddCommand(getAccountNameCmd)
}

func getAccountName() (string, error) {
	client, err := omglol.NewClient(viper.GetString("email"), viper.GetString("apikey"), endpoint)
	cobra.CheckErr(err)
	account, err := client.GetAccountName()
	return *account, err
}
