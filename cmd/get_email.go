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

	"github.com/mcornick/omglol-client-go/omglol"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var getEmailCmd = &cobra.Command{
	Use:     "email",
	Aliases: []string{"emails"},
	Short:   "Get email forwarding address(es)",
	Long:    "Gets your email forwarding address(es).",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := getEmail()
		handleAPIError(err)
		if len(result) > 1 {
			fmt.Println(strings.Join(result, ","))
		} else {
			fmt.Println(result[0])
		}
	},
}

func init() {
	getCmd.AddCommand(getEmailCmd)
}

func getEmail() ([]string, error) {
	client, err := omglol.NewClient(viper.GetString("email"), viper.GetString("apikey"), endpoint)
	cobra.CheckErr(err)
	account, err := client.GetEmails(viper.GetString("address"))
	return account, err
}
