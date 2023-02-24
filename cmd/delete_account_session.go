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

var deleteAccountSessionCmd = &cobra.Command{
	Use:   "session",
	Short: "Delete a session",
	Long: `Deletes an active session, logging it out.

Specify the session ID with the --id flag.

Note that you won't be asked to confirm deletion.
Be sure you know what you're doing.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		type Result struct {
			Request  responseRequest `json:"request"`
			Response struct {
				Message string `json:"message"`
			} `json:"response"`
		}
		var result Result
		body := callAPIWithParams(
			http.MethodDelete,
			"/account/"+viper.GetString("email")+"/sessions/"+idFlag,
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
	deleteAccountSessionCmd.Flags().StringVarP(
		&idFlag,
		"id",
		"i",
		"",
		"ID of the session to delete",
	)
	deleteAccountSessionCmd.MarkFlagRequired("id")
	deleteAccountCmd.AddCommand(deleteAccountSessionCmd)
}
