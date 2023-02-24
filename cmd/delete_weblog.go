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

var deleteWeblogCmd = &cobra.Command{
	Use:   "weblog [id]",
	Short: "Delete a weblog entry",
	Long: `Deletes a weblog entry.

Note that you won't be asked to confirm deletion.
Be sure you know what you're doing.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		type output struct {
			Request  resultRequest `json:"request"`
			Response struct {
				Message string `json:"message"`
			} `json:"response"`
		}
		var result output
		body := callAPIWithParams(
			http.MethodDelete,
			"/address/"+viper.GetString("address")+"/weblog/delete/"+args[0],
			nil,
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
	deleteCmd.AddCommand(deleteWeblogCmd)
}
