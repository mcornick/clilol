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

type getAccountSettingsOutput struct {
	Request  resultRequest `json:"request"`
	Response struct {
		Message  string `json:"message"`
		Settings struct {
			Owner         string `json:"owner"`
			Communication string `json:"communication"`
			DateFormat    string `json:"date_format"`
			WebEditor     string `json:"web_editor"`
		} `json:"settings"`
	} `json:"response"`
}

var getAccountSettingsCmd = &cobra.Command{
	Use:   "settings",
	Short: "Get your account settings",
	Long:  "Gets the settings on your account.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		var result getAccountSettingsOutput
		result, err := getAccountSettings()
		if err != nil {
			return err
		}
		if result.Request.Success {
			fmt.Println(result.Response.Message)
			fmt.Printf("Owner: %s\n", result.Response.Settings.Owner)
			fmt.Printf("Communication: %s\n", result.Response.Settings.Communication)
			fmt.Printf("Date Format: %s\n", result.Response.Settings.DateFormat)
			fmt.Printf("Web Editor: %s\n", result.Response.Settings.WebEditor)
		} else {
			return fmt.Errorf("%s", result.Response.Message)
		}
		return nil
	},
}

func init() {
	getAccountCmd.AddCommand(getAccountSettingsCmd)
}

func getAccountSettings() (getAccountSettingsOutput, error) {
	var result getAccountSettingsOutput
	body, err := callAPIWithParams(
		http.MethodGet,
		"/account/"+viper.GetString("email")+"/settings",
		nil,
		true,
	)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}
