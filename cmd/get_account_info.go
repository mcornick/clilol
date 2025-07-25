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

type getAccountInfoOutputResponse struct {
	Message string `json:"message"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Created struct {
		UnixEpochTime int64     `json:"unix_epoch_time"`
		ISO8601Time   time.Time `json:"iso_8601_time"`
		RFC2822Time   string    `json:"rfc_2822_time"`
		RelativeTime  string    `json:"relative_time"`
	} `json:"created"`
	Settings struct {
		Communication string `json:"communication"`
	} `json:"settings"`
}
type getAccountInfoOutput struct {
	Request  resultRequest                `json:"request"`
	Response getAccountInfoOutputResponse `json:"response"`
}

var getAccountInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info about your account",
	Long:  "Gets information about your account",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		var result getAccountInfoOutput
		result, err := getAccountInfo()
		if err != nil {
			return nil
		}
		if result.Request.Success {
			fmt.Println(result.Response.Message)
			fmt.Printf("%s (%s)\n", result.Response.Name, result.Response.Email)
			fmt.Printf("Created %s\n", result.Response.Created.RelativeTime)
			fmt.Printf("Communication: %s\n", result.Response.Settings.Communication)
		} else {
			return fmt.Errorf("%s", result.Response.Message)
		}
		return nil
	},
}

func init() {
	getAccountCmd.AddCommand(getAccountInfoCmd)
}

func getAccountInfo() (getAccountInfoOutput, error) {
	var result getAccountInfoOutput
	body, err := callAPIWithParams(
		http.MethodGet,
		"/account/"+viper.GetString("email")+"/info",
		nil,
		true,
	)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}
