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

	"github.com/mcornick/omglol-client-go/omglol"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var getPURLCmd = &cobra.Command{
	Use:   "purl [name]",
	Short: "Get a PURL",
	Long: `Gets a PURL by name.

The address can be specified with the --address flag. If not set,
it defaults to your own address.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		result, err := getPURL(addressFlag, args[0])
		handleAPIError(err)
		listed := "listed"
		if !result.Listed {
			listed = "unlisted"
		}
		fmt.Printf(
			"%s: %s (%s)\n",
			result.Name,
			result.URL,
			listed,
		)
	},
}

func init() {
	getPURLCmd.Flags().StringVarP(
		&addressFlag,
		"address",
		"a",
		"",
		"address whose PURL to get",
	)
	getCmd.AddCommand(getPURLCmd)
}

func getPURL(address string, name string) (omglol.PersistentURL, error) {
	if address == "" {
		address = viper.GetString("address")
	}
	client, err := omglol.NewClient(viper.GetString("email"), viper.GetString("apikey"), endpoint)
	cobra.CheckErr(err)
	purl, err := client.GetPersistentURL(address, name)
	return *purl, err
}
