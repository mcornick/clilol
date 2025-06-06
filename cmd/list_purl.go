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

var listPURLCmd = &cobra.Command{
	Use:     "purls",
	Aliases: []string{"purl"},
	Short:   "List all PURLs",
	Long: `Lists all PURLs for a user.

The address can be specified with the --address flag. If not set,
it defaults to your own address.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := listPURL(addressFlag)
		handleAPIError(err)
		for _, purl := range result {
			listed := "listed"
			if !purl.Listed {
				listed = "unlisted"
			}
			fmt.Printf(
				"%s: %s (%s)\n",
				purl.Name,
				purl.URL,
				listed,
			)
		}
	},
}

func init() {
	listPURLCmd.Flags().StringVarP(
		&addressFlag,
		"address",
		"a",
		"",
		"address whose PURLs to get",
	)
	listCmd.AddCommand(listPURLCmd)
}

func listPURL(address string) ([]omglol.PersistentURL, error) {
	if address == "" {
		address = viper.GetString("address")
	}
	client, err := omglol.NewClient(viper.GetString("email"), viper.GetString("apikey"), endpoint)
	cobra.CheckErr(err)
	purls, err := client.ListPersistentURLs(address)
	return *purls, err
}
