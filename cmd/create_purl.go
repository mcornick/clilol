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

type createPURLInput struct {
	Name   string `json:"name"`
	URL    string `json:"url"`
	Listed bool   `json:"listed,omitempty"`
}
type createPURLOutput struct {
	Request  resultRequest `json:"request"`
	Response struct {
		Message string `json:"message"`
		Name    string `json:"name"`
		URL     string `json:"url"`
	} `json:"response"`
}

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
		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := createPURL(args[0], args[1], createPURLListed)
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

func createPURL(name string, url string, listed bool) (createPURLOutput, error) {
	var result createPURLOutput
	purl := createPURLInput{name, url, listed}
	body, err := callAPIWithParams(
		http.MethodPost,
		"/address/"+viper.GetString("address")+"/purl",
		purl,
		true,
	)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}
