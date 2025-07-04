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
)

type getAddressExpirationOutput struct {
	Request  resultRequest `json:"request"`
	Response struct {
		Message string `json:"message"`
		Expired bool   `json:"expired"`
	} `json:"response"`
}

var getAddressExpirationCmd = &cobra.Command{
	Use:   "expiration [address]",
	Short: "Get address expiration",
	Long:  "Gets the expiration of an address.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := getAddressExpiration(args[0])
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
	getAddressCmd.AddCommand(getAddressExpirationCmd)
}

func getAddressExpiration(address string) (getAddressExpirationOutput, error) {
	var result getAddressExpirationOutput
	body, err := callAPIWithParams(
		http.MethodGet,
		"/address/"+address+"/expiration",
		nil,
		false,
	)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}
