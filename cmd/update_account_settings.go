// SPDX-License-Identifier: MPL-2.0
//
// Copyright (c) 2023 Mark Cornick
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type updateAccountSettingsInput struct {
	Communication string `json:"communication,omitempty"`
	DateFormat    string `json:"date_format,omitempty"`
	WebEditor     string `json:"web_editor,omitempty"`
}

type updateAccountSettingsOutput struct {
	Request  resultRequest `json:"request"`
	Response struct {
		Message string `json:"message"`
	} `json:"response"`
}

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
			validateConfig()
			result, err := updateAccountSettings(accountUpdateSettingsCommunication, accountUpdateSettingsDateFormat, accountUpdateSettingsWebEditor)
			cobra.CheckErr(err)
			if result.Request.Success {
				fmt.Println(result.Response.Message)
			} else {
				cobra.CheckErr(fmt.Errorf(result.Response.Message))
			}
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

func updateAccountSettings(communication string, dateFormat string, webEditor string) (updateAccountSettingsOutput, error) {
	var result updateAccountSettingsOutput
	account := updateAccountSettingsInput{communication, dateFormat, webEditor}
	body := callAPIWithParams(
		http.MethodPost,
		"/account/"+viper.GetString("email")+"/settings",
		account,
		true,
	)
	err := json.Unmarshal(body, &result)
	return result, err
}
