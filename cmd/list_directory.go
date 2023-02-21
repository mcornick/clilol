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

var listDirectoryCmd = &cobra.Command{
	Use:   "directory",
	Short: "list the address directory",
	Long:  "Lists the omg.lol address directory.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		type Result struct {
			Request struct {
				StatusCode int  `json:"status_code"`
				Success    bool `json:"success"`
			} `json:"request"`
			Response struct {
				Message   string   `json:"message"`
				URL       string   `json:"url"`
				Directory []string `json:"directory"`
			} `json:"response"`
		}
		var result Result
		body := callAPIWithJSON(http.MethodGet, "/directory", nil, false)
		err := json.Unmarshal(body, &result)
		cobra.CheckErr(err)
		if !silent {
			if !wantJson {
				if result.Request.Success {
					fmt.Printf("%s\n\n", result.Response.Message)
					for _, address := range result.Response.Directory {
						idnaProfile := idna.New()
						decoded, err := idnaProfile.ToUnicode(address)
						cobra.CheckErr(err)
						fmt.Println(decoded)
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
	listCmd.AddCommand(listDirectoryCmd)
}
