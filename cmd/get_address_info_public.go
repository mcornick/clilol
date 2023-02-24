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
)

type getAddressInfoPublicOutput struct {
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
			Message string `json:"message"`
			Expired bool   `json:"expired"`
		} `json:"expiration"`
		Verification struct {
			Message  string `json:"message"`
			Verified bool   `json:"verified"`
		} `json:"verification"`
	} `json:"response"`
}

var getAddressInfoPublicCmd = &cobra.Command{
	Use:   "public [address]",
	Short: "Get public information about an address",
	Long:  "Gets public information about an address.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var result getAddressInfoPublicOutput
		result, err := getAddressInfoPublic(args[0])
		cobra.CheckErr(err)
		if result.Request.Success {
			fmt.Println(result.Response.Registration.Message)
			fmt.Println(result.Response.Expiration.Message)
			fmt.Println(result.Response.Verification.Message)
		} else {
			cobra.CheckErr(fmt.Errorf(result.Response.Message))
		}
	},
}

func init() {
	getAddressInfoCmd.AddCommand(getAddressInfoPublicCmd)
}

func getAddressInfoPublic(address string) (getAddressInfoPublicOutput, error) {
	var result getAddressInfoPublicOutput
	body := callAPIWithParams(
		http.MethodGet,
		"/address/"+address+"/info",
		nil,
		false,
	)
	err := json.Unmarshal(body, &result)
	return result, err
}
