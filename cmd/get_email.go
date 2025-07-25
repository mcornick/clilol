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

type getEmailOutput struct {
	Request  resultRequest `json:"request"`
	Response struct {
		Message           string   `json:"message"`
		DestinationString string   `json:"destination_string"`
		DestinationArray  []string `json:"destination_array"`
		Address           string   `json:"address"`
		EmailAddress      string   `json:"email_address"`
	} `json:"response"`
}

var getEmailCmd = &cobra.Command{
	Use:   "email",
	Short: "Get email forwarding address(es)",
	Long:  "Gets your email forwarding address(es).",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := getEmail()
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
	getCmd.AddCommand(getEmailCmd)
}

func getEmail() (getEmailOutput, error) {
	var result getEmailOutput
	body, err := callAPIWithParams(
		http.MethodGet,
		"/address/"+viper.GetString("address")+"/email",
		nil,
		true,
	)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}
