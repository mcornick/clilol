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

type listPURLOutput struct {
	Request  resultRequest `json:"request"`
	Response struct {
		Message string `json:"message"`
		PURLs   []struct {
			Name    string `json:"name"`
			URL     string `json:"url"`
			Counter int    `json:"counter"`
		} `json:"purls"`
	} `json:"response"`
}

var listPURLCmd = &cobra.Command{
	Use:   "purl",
	Short: "List all PURLs",
	Long: `Lists all PURLs for a user.

The address can be specified with the --address flag. If not set,
it defaults to your own address.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := listPURL(addressFlag)
		cobra.CheckErr(err)
		if result.Request.Success {
			for _, purl := range result.Response.PURLs {
				fmt.Printf(
					"%s: %s (%d hits)\n",
					purl.Name,
					purl.URL,
					purl.Counter,
				)
			}
		} else {
			cobra.CheckErr(fmt.Errorf(result.Response.Message))
		}
	},
}

func init() {
	listPURLCmd.Flags().StringVarP(
		&addressFlag,
		"address",
		"a",
		"",
		"address whose PURLs to get",
	)
	listCmd.AddCommand(listPURLCmd)
}

func listPURL(address string) (listPURLOutput, error) {
	var result listPURLOutput
	if address == "" {
		address = viper.GetString("address")
	}
	body := callAPIWithParams(
		http.MethodGet,
		"/address/"+address+"/purls",
		nil,
		true,
	)
	err := json.Unmarshal(body, &result)
	return result, err
}
