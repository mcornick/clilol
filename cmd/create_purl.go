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

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	createPURLURL    string
	createPURLListed bool
	createPURLCmd    = &cobra.Command{
		Use:   "purl",
		Short: "Create a PURL",
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
			purl := Input{nameFlag, createPURLURL, createPURLListed}
			body := callAPIWithParams(
				http.MethodPost,
				"/address/"+viper.GetString("address")+"/purl",
				purl,
				true,
			)
			err := json.Unmarshal(body, &result)
			checkError(err)
			if result.Request.Success {
				log.Info(result.Response.Message)
			} else {
				checkError(fmt.Errorf(result.Response.Message))
			}
		},
	}
)

func init() {
	createPURLCmd.Flags().StringVarP(
		&nameFlag,
		"name",
		"n",
		"",
		"Name of the PURL",
	)
	createPURLCmd.Flags().StringVarP(
		&createPURLURL,
		"url",
		"u",
		"",
		"URL to redirect to",
	)
	createPURLCmd.Flags().BoolVarP(
		&createPURLListed,
		"listed",
		"l",
		false,
		"Create as listed (default false)",
	)
	createCmd.AddCommand(createPURLCmd)
}
