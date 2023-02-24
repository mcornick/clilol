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
	createPURLListed bool
	createPURLCmd    = &cobra.Command{
		Use:   "purl [name] [url]",
		Short: "Create a PURL",
		Long: `Creates a PURL.

The PURL will be created as unlisted by default. To create a listed
PURL, use the --listed flag.
`,
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			type input struct {
				Name   string `json:"name"`
				URL    string `json:"url"`
				Listed bool   `json:"listed,omitempty"`
			}
			type output struct {
				Request  resultRequest `json:"request"`
				Response struct {
					Message string `json:"message"`
					Name    string `json:"name"`
					URL     string `json:"url"`
				} `json:"response"`
			}
			var result output
			purl := input{args[0], args[1], createPURLListed}
			body := callAPIWithParams(
				http.MethodPost,
				"/address/"+viper.GetString("address")+"/purl",
				purl,
				true,
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
)

func init() {
	createPURLCmd.Flags().BoolVarP(
		&createPURLListed,
		"listed",
		"l",
		false,
		"create as listed (default false)",
	)
	createCmd.AddCommand(createPURLCmd)
}
