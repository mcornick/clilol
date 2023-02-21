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

var listPURLCmd = &cobra.Command{
	Use:   "purl",
	Short: "list all PURLs",
	Long: `Lists all PURLs for a user.

The address can be specified with the --address flag. If not set,
it defaults to your own address.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		type Result struct {
			Request struct {
				StatusCode int  `json:"status_code"`
				Success    bool `json:"success"`
			} `json:"request"`
			Response struct {
				Message string `json:"message"`
				PURLs   []struct {
					Name    string `json:"name"`
					URL     string `json:"url"`
					Counter int    `json:"counter"`
				} `json:"purls"`
			} `json:"response"`
		}
		var result Result
		body := callAPIWithJSON(
			http.MethodGet,
			"/address/"+address+"/purls",
			nil,
			true,
		)
		err := json.Unmarshal(body, &result)
		cobra.CheckErr(err)
		if !silent {
			if !wantJson {
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
			} else {
				fmt.Println(string(body))
			}
		}
	},
}

func init() {
	listPURLCmd.Flags().StringVarP(
		&address,
		"address",
		"a",
		viper.GetString("address"),
		"address whose PURLs to get",
	)
	listCmd.AddCommand(listPURLCmd)
}
