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

	"github.com/ejstreet/omglol-client-go/omglol"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listAccountSessionsCmd = &cobra.Command{
	Use:     "sessions",
	Aliases: []string{"session"},
	Short:   "List sessions on your account",
	Long:    "Lists active sessions on your account.",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		validateConfig()
		sessions, err := listAccountSessions()
		handleAPIError(err)
		for _, session := range sessions {
			fmt.Printf("\n%s\n", session.SessionID)
			fmt.Println(session.UserAgent)
			fmt.Println(session.CreatedIP)
			fmt.Println(time.Unix(session.CreatedOn, 0))
		}
	},
}

func init() {
	listAccountCmd.AddCommand(listAccountSessionsCmd)
}

func listAccountSessions() ([]omglol.ActiveSession, error) {
	client, err := omglol.NewClient(viper.GetString("email"), viper.GetString("apikey"), endpoint)
	cobra.CheckErr(err)
	sessions, err := client.GetActiveSessions()
	return *sessions, err
}
