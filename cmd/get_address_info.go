// SPDX-License-Identifier: MPL-2.0
//
// Copyright (c) 2023 Mark Cornick
//
// This Source Code Form is subject to the terms of the Mozilla Pu
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

type getAddressInfoOutput struct {
	Request  resultRequest `json:"request"`
	Response struct {
		Address      string `json:"address"`
		Message      string `json:"message"`
		Registration struct {
			Message       string    `json:"message"`
			UnixEpochTime int64     `json:"unix_epoch_time"`
			ISO8601Time   time.Time `json:"iso_8601_time"`
			RFC2822Time   string    `json:"rfc_2822_time"`
			RelativeTime  string    `json:"relative_time"`
		} `json:"registration"`
		Expiration struct {
			Message       string    `json:"message"`
			Expired       bool      `json:"expired"`
			WillExpire    bool      `json:"will_expire"`
			UnixEpochTime int64     `json:"unix_epoch_time"`
			ISO8601Time   time.Time `json:"iso_8601_time"`
			RFC2822Time   string    `json:"rfc_2822_time"`
			RelativeTime  string    `json:"relative_time"`
		} `json:"expiration"`
		Verification struct {
			Message  string `json:"message"`
			Verified bool   `json:"verified"`
		} `json:"verification"`
		Owner string `json:"owner"`
	} `json:"response"`
}

var getAddressInfoCmd = &cobra.Command{
	Use:   "info [name]",
	Short: "Get information about an address",
	Long:  "Gets information about an address.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := getAddressInfo(args[0])
		if err != nil {
			return err
		}
		if result.Request.Success {
			fmt.Println(result.Response.Registration.Message)
			fmt.Println(result.Response.Expiration.Message)
			fmt.Println(result.Response.Verification.Message)
		} else {
			return fmt.Errorf("%s", result.Response.Message)
		}
		return nil
	},
}

func init() {
	getAddressCmd.AddCommand(getAddressInfoCmd)
}

func getAddressInfo(name string) (getAddressInfoOutput, error) {
	var result getAddressInfoOutput
	body, err := callAPIWithParams(
		http.MethodGet,
		"/address/"+name+"/info",
		nil,
		true,
	)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}
