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

type deleteAccountSessionOutput struct {
	Request  resultRequest `json:"request"`
	Response struct {
		Message string `json:"message"`
	} `json:"response"`
}

var deleteAccountSessionCmd = &cobra.Command{
	Use:   "session [id]",
	Short: "Delete a session",
	Long: `Deletes an active session, logging it out.

Note that you won't be asked to confirm deletion.
Be sure you know what you're doing.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		result, err := deleteAccountSession(args[0])
		cobra.CheckErr(err)
		if result.Request.Success {
			fmt.Println(result.Response.Message)
		} else {
			cobra.CheckErr(fmt.Errorf(result.Response.Message))
		}
	},
}

func init() {
	deleteAccountCmd.AddCommand(deleteAccountSessionCmd)
}

func deleteAccountSession(id string) (deleteAccountSessionOutput, error) {
	var result deleteAccountSessionOutput
	body := callAPIWithParams(
		http.MethodDelete,
		"/account/"+viper.GetString("email")+"/sessions/"+id,
		nil,
		true,
	)
	err := json.Unmarshal(body, &result)
	return result, err
}
