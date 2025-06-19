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
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type listAccountSessionsOutput struct {
	Request  resultRequest `json:"request"`
	Response []struct {
		SessionID string `json:"session_id"`
		UserAgent string `json:"user_agent"`
		CreatedIP string `json:"created_ip"`
		CreatedOn int64  `json:"created_on"`
		ExpiresOn int64  `json:"expires_on"`
		UpdatedOn int64  `json:"updated_on"`
	} `json:"response"`
}

var listAccountSessionsCmd = &cobra.Command{
	Use:     "sessions",
	Aliases: []string{"session"},
	Short:   "List your sessions",
	Long:    "Lists the active sessions on your account.",
	Args:    cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := listAccountSessions()
		if err != nil {
			return err
		}
		if result.Request.Success {
			for _, session := range result.Response {
				fmt.Printf("\n%s\n", session.SessionID)
				fmt.Println(session.UserAgent)
				fmt.Println(session.CreatedIP)
				fmt.Println(time.Unix(session.CreatedOn, 0))
			}
		} else {
			return fmt.Errorf("%d", result.Request.StatusCode)
		}
		return nil
	},
}

func init() {
	listAccountCmd.AddCommand(listAccountSessionsCmd)
}

func listAccountSessions() (listAccountSessionsOutput, error) {
	var result listAccountSessionsOutput
	body := callAPIWithParams(
		http.MethodGet,
		"/account/"+viper.GetString("email")+"/sessions",
		nil,
		true,
	)
	err := json.Unmarshal(body, &result)
	return result, err
}
