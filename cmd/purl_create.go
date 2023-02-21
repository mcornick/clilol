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
	purlCreateURL    string
	purlCreateListed bool
	purlCreateCmd    = &cobra.Command{
		Use:   "create",
		Short: "create a PURL",
		Long: `Creates a PURL.

Specify the PURL name with the --name flag, and the URL with the
--url flag.`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			type Input struct {
				Name   string `json:"name"`
				URL    string `json:"url"`
				Listed bool   `json:"listed,omitempty"`
			}
			type Result struct {
				Request struct {
					StatusCode int  `json:"status_code"`
					Success    bool `json:"success"`
				} `json:"request"`
				Response struct {
					Message string `json:"message"`
					Name    string `json:"name"`
					URL     string `json:"url"`
				} `json:"response"`
			}
			var result Result
			purl := Input{name, purlCreateURL, purlCreateListed}
			body := callAPI(
				http.MethodPost,
				"/address/"+viper.GetString("username")+"/purl",
				purl,
				true,
			)
			err := json.Unmarshal(body, &result)
			cobra.CheckErr(err)
			if !silent {
				if !wantJson {
					if result.Request.Success {
						fmt.Println(result.Response.Message)
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
	purlCreateCmd.Flags().StringVarP(
		&name,
		"name",
		"n",
		"",
		"Name of the PURL",
	)
	purlCreateCmd.Flags().StringVarP(
		&purlCreateURL,
		"url",
		"u",
		"",
		"URL to redirect to",
	)
	purlCreateCmd.Flags().BoolVarP(
		&purlCreateListed,
		"listed",
		"l",
		false,
		"Create as listed (default false)",
	)
	purlCmd.AddCommand(purlCreateCmd)
}
