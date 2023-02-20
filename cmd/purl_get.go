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
	purlGetName string
	purlGetCmd  = &cobra.Command{
		Use:   "get",
		Short: "get a PURL",
		Long: `Gets a PURL by name.

Specify the name with the --name flag.

The username can be specified with the --username flag. If not set,
it defaults to your own username.`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			type Result struct {
				Request struct {
					StatusCode int  `json:"status_code"`
					Success    bool `json:"success"`
				} `json:"request"`
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
			body := callAPI(
				http.MethodGet,
				"/address/"+username+"/purl/"+purlGetName,
				nil,
				true,
			)
			err := json.Unmarshal(body, &result)
			cobra.CheckErr(err)
			if !silent {
				if !wantJson {
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
				} else {
					fmt.Println(string(body))
				}
			}
		},
	}
)

func init() {
	purlGetCmd.Flags().StringVarP(
		&username,
		"username",
		"u",
		viper.GetString("username"),
		"username whose PURL to get",
	)
	purlGetCmd.Flags().StringVarP(
		&purlGetName,
		"name",
		"n",
		"",
		"name of the PURL to get",
	)
	purlCmd.AddCommand(purlGetCmd)
}
