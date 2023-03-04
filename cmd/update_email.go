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
	"strings"

	"github.com/ejstreet/omglol-client-go/omglol"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var updateEmailCmd = &cobra.Command{
	Use:     "email [address]",
	Aliases: []string{"emails"},
	Short:   "set email forwarding address(es)",
	Long: `Sets your email forwarding address(es).
	
To specify multiple addresses, separate them with commas.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		validateConfig()
		destination := strings.Split(args[0], ",")
		err := updateEmail(destination)
		handleAPIError(err)
		fmt.Println("Email forwarding updated.")
	},
}

func init() {
	updateCmd.AddCommand(updateEmailCmd)
}

func updateEmail(destination []string) error {
	client, err := omglol.NewClient(viper.GetString("email"), viper.GetString("apikey"), endpoint)
	cobra.CheckErr(err)
	err = client.SetEmails(viper.GetString("address"), destination)
	return err
}
