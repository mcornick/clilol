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

var (
	getPURLName string
	getPURLCmd  = &cobra.Command{
		Use:   "purl",
		Short: "Get a PURL",
		Long: `Gets a PURL by name.

Specify the name with the --name flag.

The address can be specified with the --address flag. If not set,
it defaults to your own address.`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			type Result struct {
				Request  responseRequest `json:"request"`
				Response struct {
					Message string `json:"message"`
					PURL    struct {
						Name    string `json:"name"`
						URL     string `json:"url"`
						Counter int    `json:"counter"`
					} `json:"purl"`
				} `json:"response"`
			}
			var result Result
			if addressFlag == "" {
				addressFlag = viper.GetString("address")
			}
			body := callAPIWithParams(
				http.MethodGet,
				"/address/"+addressFlag+"/purl/"+getPURLName,
				nil,
				true,
			)
			err := json.Unmarshal(body, &result)
			cobra.CheckErr(err)
			if result.Request.Success {
				fmt.Printf(
					"%s: %s (%d hits)\n",
					result.Response.PURL.Name,
					result.Response.PURL.URL,
					result.Response.PURL.Counter,
				)
			} else {
				cobra.CheckErr(fmt.Errorf(result.Response.Message))
			}
		},
	}
)

func init() {
	getPURLCmd.Flags().StringVarP(
		&addressFlag,
		"address",
		"a",
		"",
		"address whose PURL to get",
	)
	getPURLCmd.Flags().StringVarP(
		&getPURLName,
		"name",
		"n",
		"",
		"name of the PURL to get",
	)
	getPURLCmd.MarkFlagRequired("name")
	getCmd.AddCommand(getPURLCmd)
}
