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

var deleteAccountSessionCmd = &cobra.Command{
	Use:   "session [id]",
	Short: "Delete a session",
	Long: `Deletes an active session, logging it out.

Note that you won't be asked to confirm deletion.
Be sure you know what you're doing.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		err := deleteAccountSession(id)
		handleAPIError(err)
		fmt.Printf("Session %s deleted.\n", id)
	},
}

func init() {
	deleteAccountCmd.AddCommand(deleteAccountSessionCmd)
}

func deleteAccountSession(id string) error {
	client, err := omglol.NewClient(viper.GetString("email"), viper.GetString("apikey"), endpoint)
	cobra.CheckErr(err)
	err = client.DeleteActiveSession(id)
	return err
}
