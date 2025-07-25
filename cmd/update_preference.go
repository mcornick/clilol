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

type updatePreferenceInput struct {
	Item  string `json:"item"`
	Value string `json:"value"`
}
type updatePreferenceOutput struct {
	Response struct {
		Message string `json:"message"`
		Item    string `json:"item"`
		Value   string `json:"value"`
	} `json:"response"`
	Request resultRequest `json:"request"`
}

var updatePreferenceCmd = &cobra.Command{
	Use:   "preference [item] [value]",
	Short: "set a preference",
	Long:  "Sets omg.lol preferences.",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := updatePreference(args[0], args[1])
		if err != nil {
			return err
		}
		if result.Request.Success {
			fmt.Println(result.Response.Message)
		} else {
			return fmt.Errorf("%s", result.Response.Message)
		}
		return nil
	},
}

func init() {
	updateCmd.AddCommand(updatePreferenceCmd)
}

func updatePreference(item string, value string) (updatePreferenceOutput, error) {
	var result updatePreferenceOutput
	pref := updatePreferenceInput{item, value}
	body, err := callAPIWithParams(
		http.MethodPost,
		"/preferences/"+viper.GetString("address"),
		pref,
		true,
	)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}
