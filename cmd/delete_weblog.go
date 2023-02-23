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

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var deleteWeblogCmd = &cobra.Command{
	Use:   "weblog",
	Short: "Delete a weblog entry",
	Long: `Deletes a weblog entry.

Specify the weblog entry ID with the --id flag.

Note that you won't be asked to confirm deletion.
Be sure you know what you're doing.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
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
		body := callAPIWithParams(
			http.MethodDelete,
			"/address/"+viper.GetString("address")+"/weblog/delete/"+idFlag,
			nil,
			true,
		)
		err := json.Unmarshal(body, &result)
		checkError(err)
		if result.Request.Success {
			log.Info(result.Response.Message)
		} else {
			checkError(fmt.Errorf(result.Response.Message))
		}
	},
}

func init() {
	deleteWeblogCmd.Flags().StringVarP(
		&idFlag,
		"id",
		"i",
		"",
		"ID of the weblog entry to delete",
	)
	deleteCmd.AddCommand(deleteWeblogCmd)
}
