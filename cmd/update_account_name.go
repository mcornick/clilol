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

var updateAccountNameCmd = &cobra.Command{
	Use:   "name [name]",
	Short: "set the name on your account",
	Long:  "Sets the name on your account.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := updateAccountName(args[0])
		handleAPIError(err)
		fmt.Println("Name updated.")
	},
}

func init() {
	updateAccountCmd.AddCommand(updateAccountNameCmd)
}

func updateAccountName(name string) error {
	client, err := omglol.NewClient(viper.GetString("email"), viper.GetString("apikey"), endpoint)
	cobra.CheckErr(err)
	err = client.SetAccountName(name)
	return err
}
