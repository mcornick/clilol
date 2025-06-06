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
	"golang.org/x/net/idna"
)

var listDirectoryCmd = &cobra.Command{
	Use:   "directory",
	Short: "List the address directory",
	Long:  "Lists the omg.lol address directory.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		validateConfig()
		result, err := listDirectory()
		handleAPIError(err)
		for _, address := range result.Directory {
			idnaProfile := idna.New()
			decoded, err := idnaProfile.ToUnicode(address)
			cobra.CheckErr(err)
			fmt.Println(decoded)
		}
	},
}

func init() {
	listCmd.AddCommand(listDirectoryCmd)
}

func listDirectory() (*omglol.AddressDirectory, error) {
	client, err := omglol.NewClient(viper.GetString("email"), viper.GetString("apikey"), endpoint)
	cobra.CheckErr(err)
	directory, err := client.GetAddressDirectory()
	return directory, err
}
