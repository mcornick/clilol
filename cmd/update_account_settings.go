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

var (
	accountUpdateSettingsCommunication string
	accountUpdateSettingsDateFormat    string
	accountUpdateSettingsWebEditor     string
	updateAccountSettingsCmd           = &cobra.Command{
		Use:   "settings",
		Short: "set the settings on your account",
		Long: `Sets the settings on your account.

Specify the new settings with the flags listed below.`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			err := updateAccountSettings(accountUpdateSettingsCommunication, accountUpdateSettingsDateFormat, accountUpdateSettingsWebEditor)
			handleAPIError(err)
			fmt.Println("Settings updated.")
		},
	}
)

func init() {
	updateAccountSettingsCmd.Flags().StringVarP(
		&accountUpdateSettingsCommunication,
		"communication",
		"c",
		"",
		"communication preference",
	)
	updateAccountSettingsCmd.Flags().StringVarP(
		&accountUpdateSettingsDateFormat,
		"date-format",
		"d",
		"",
		"date format preference",
	)
	updateAccountSettingsCmd.Flags().StringVarP(
		&accountUpdateSettingsWebEditor,
		"web-editor",
		"w",
		"",
		"web editor preference",
	)
	updateAccountCmd.AddCommand(updateAccountSettingsCmd)
}

func updateAccountSettings(communication string, dateFormat string, webEditor string) error {
	client, err := omglol.NewClient(viper.GetString("email"), viper.GetString("apikey"), endpoint)
	cobra.CheckErr(err)
	settings := map[string]string{
		"communication": communication,
		"date_format":   dateFormat,
		"web_editor":    webEditor,
	}
	err = client.SetAccountSettings(settings)
	return err
}
