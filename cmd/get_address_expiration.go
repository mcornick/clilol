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

var getAddressExpirationCmd = &cobra.Command{
	Use:   "expiration",
	Short: "Get address expiration",
	Long: `Gets the expiration of an address.
	
Specify the address with the --address flag.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		type Result struct {
			Request  responseRequest `json:"request"`
			Response struct {
				Message string `json:"message"`
				Expired bool   `json:"expired"`
			} `json:"response"`
		}
		var result Result
		if addressFlag == "" {
			addressFlag = viper.GetString("address")
		}
		body := callAPIWithParams(
			http.MethodGet,
			"/address/"+addressFlag+"/expiration",
			nil,
			false,
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
	getAddressExpirationCmd.Flags().StringVarP(
		&addressFlag,
		"address",
		"a",
		"",
		"address whose expiration to get",
	)
	err := getAddressExpirationCmd.MarkFlagRequired("address")
	cobra.CheckErr(err)
	getAddressCmd.AddCommand(getAddressExpirationCmd)
}
