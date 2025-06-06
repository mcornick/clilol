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
	"time"

	"github.com/mcornick/omglol-client-go/omglol"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listPasteCmd = &cobra.Command{
	Use:     "pastes",
	Aliases: []string{"paste"},
	Short:   "List pastes",
	Long: `Lists pastes for a user.

The address can be specified with the --address flag. If not set,
it defaults to your own address.

Unlisted pastes are only included when the --address flag is set to
your own address.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		validateConfig()
		result, err := listPaste(addressFlag)
		handleAPIError(err)
		for _, paste := range result {
			if paste.ModifiedOn != nil {
				fmt.Printf(
					"%s modified on %s\n",
					paste.Title,
					time.Unix(*paste.ModifiedOn, 0),
				)
			}
		}
	},
}

func init() {
	listPasteCmd.Flags().StringVarP(
		&addressFlag,
		"address",
		"a",
		"",
		"address whose pastes to list",
	)
	listCmd.AddCommand(listPasteCmd)
}

func listPaste(address string) ([]omglol.Paste, error) {
	if address == "" {
		address = viper.GetString("address")
	}
	client, err := omglol.NewClient(viper.GetString("email"), viper.GetString("apikey"), endpoint)
	cobra.CheckErr(err)
	pastes, err := client.ListPastes(address)
	return *pastes, err
}
