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

var (
	accountSetSettingsCommunication string
	accountSetSettingsDateFormat    string
	accountSetSettingsWebEditor     string
	updateAccountSettingsCmd        = &cobra.Command{
		Use:   "settings",
		Short: "set the settings on your account",
		Long: `Sets the settings on your account.

Specify the new settings with the flags listed below.`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			type Input struct {
				Communication string `json:"communication,omitempty"`
				DateFormat    string `json:"date_format,omitempty"`
				WebEditor     string `json:"web_editor,omitempty"`
			}
			type Result struct {
				Request struct {
					StatusCode int  `json:"status_code"`
					Success    bool `json:"success"`
				} `json:"request"`
				Response struct {
					Message string `json:"message"`
				} `json:"response"`
			}
			var result Result
			account := Input{accountSetSettingsCommunication, accountSetSettingsDateFormat, accountSetSettingsWebEditor}
			body := callAPIWithParams(
				http.MethodPost,
				"/account/"+viper.GetString("email")+"/settings",
				account,
				true,
			)
			err := json.Unmarshal(body, &result)
			checkError(err)
			if !jsonFlag {
				if result.Request.Success {
					logInfo(result.Response.Message)
				} else {
					checkError(fmt.Errorf(result.Response.Message))
				}
			} else {
				fmt.Println(string(body))
			}
		},
	}
)

func init() {
	updateAccountSettingsCmd.Flags().StringVarP(
		&accountSetSettingsCommunication,
		"communication",
		"c",
		"",
		"Communication preference",
	)
	updateAccountSettingsCmd.Flags().StringVarP(
		&accountSetSettingsDateFormat,
		"date-format",
		"d",
		"",
		"Date format preference",
	)
	updateAccountSettingsCmd.Flags().StringVarP(
		&accountSetSettingsWebEditor,
		"web-editor",
		"w",
		"",
		"Web editor preference",
	)
	updateAccountCmd.AddCommand(updateAccountSettingsCmd)
}
