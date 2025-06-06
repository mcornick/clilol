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
	"os"

	"github.com/mcornick/omglol-client-go/omglol"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	getPasteFilename string
	getPasteCmd      = &cobra.Command{
		Use:   "paste [title]",
		Short: "Get a paste",
		Long: `Gets a paste by title.

The address can be specified with the --address flag. If not set,
it defaults to your own address.`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			validateConfig()
			result, err := getPaste(addressFlag, args[0])
			handleAPIError(err)
			if getPasteFilename != "" {
				err = os.WriteFile(getPasteFilename, []byte(result.Content), 0o644)
				cobra.CheckErr(err)
			} else {
				fmt.Println(result.Content)
			}
		},
	}
)

func init() {
	getPasteCmd.Flags().StringVarP(
		&addressFlag,
		"address",
		"a",
		"",
		"address whose paste to get",
	)
	getPasteCmd.Flags().StringVarP(
		&getPasteFilename,
		"filename",
		"f",
		"",
		"file to write paste to (default stdout)",
	)
	getCmd.AddCommand(getPasteCmd)
}

func getPaste(address string, title string) (omglol.Paste, error) {
	if address == "" {
		address = viper.GetString("address")
	}
	client, err := omglol.NewClient(viper.GetString("email"), viper.GetString("apikey"), endpoint)
	cobra.CheckErr(err)
	paste, err := client.GetPaste(address, title)
	return *paste, err
}
