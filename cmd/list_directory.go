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
	"golang.org/x/net/idna"
)

type listDirectoryOutput struct {
	Request  resultRequest `json:"request"`
	Response struct {
		Message   string   `json:"message"`
		URL       string   `json:"url"`
		Directory []string `json:"directory"`
	} `json:"response"`
}

var listDirectoryCmd = &cobra.Command{
	Use:   "directory",
	Short: "List the address directory",
	Long:  "Lists the omg.lol address directory.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := listDirectory()
		if err != nil {
			return err
		}
		if result.Request.Success {
			fmt.Printf("%s\n\n", result.Response.Message)
			for _, address := range result.Response.Directory {
				idnaProfile := idna.New()
				decoded, err := idnaProfile.ToUnicode(address)
				if err != nil {
					return err
				}
				fmt.Println(decoded)
			}
		} else {
			return fmt.Errorf("%s", result.Response.Message)
		}
		return nil
	},
}

func init() {
	listCmd.AddCommand(listDirectoryCmd)
}

func listDirectory() (listDirectoryOutput, error) {
	var result listDirectoryOutput
	body, err := callAPIWithParams(http.MethodGet, "/directory", nil, false)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}
