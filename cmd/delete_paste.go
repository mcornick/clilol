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

var deletePasteCmd = &cobra.Command{
	Use:   "paste [title]",
	Short: "Delete a paste",
	Long: `Deletes a paste.

Specify the paste title with the --title flag.

Note that you won't be asked to confirm deletion.
Be sure you know what you're doing.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		validateConfig()
		title := args[0]
		err := deletePaste(title)
		handleAPIError(err)
		fmt.Printf("Paste %s deleted\n", title)
	},
}

func init() {
	deleteCmd.AddCommand(deletePasteCmd)
}

func deletePaste(title string) error {
	client, err := omglol.NewClient(viper.GetString("email"), viper.GetString("apikey"), endpoint)
	cobra.CheckErr(err)
	err = client.DeletePaste(viper.GetString("address"), title)
	return err
}
