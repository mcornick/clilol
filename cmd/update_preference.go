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

var updatePreferenceCmd = &cobra.Command{
	Use:   "preference [item] [value]",
	Short: "set a preference",
	Long:  "Sets omg.lol preferences.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		type input struct {
			Item  string `json:"item"`
			Value string `json:"value"`
		}
		type output struct {
			Request  resultRequest `json:"request"`
			Response struct {
				Message string `json:"message"`
				Item    string `json:"item"`
				Value   string `json:"value"`
			} `json:"response"`
		}
		var result output
		pref := input{args[0], args[1]}
		body := callAPIWithParams(
			http.MethodPost,
			"/preferences/"+viper.GetString("address"),
			pref,
			true,
		)
		err := json.Unmarshal(body, &result)
		cobra.CheckErr(err)
		if result.Request.Success {
			fmt.Println(result.Response.Message)
		} else {
			cobra.CheckErr(fmt.Errorf(result.Response.Message))
		}
	},
}

func init() {
	updateCmd.AddCommand(updatePreferenceCmd)
}
