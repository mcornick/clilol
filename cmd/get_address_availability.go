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
)

var getAddressAvailabilityCmd = &cobra.Command{
	Use:   "availability [address]",
	Short: "Get address availability",
	Long:  "Gets the availability of an address.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		type Result struct {
			Request  resultRequest `json:"request"`
			Response struct {
				Message      string   `json:"message"`
				Punycode     string   `json:"punycode,omitempty"`
				SeeAlso      []string `json:"see-also,omitempty"`
				Address      string   `json:"address"`
				Available    bool     `json:"available"`
				Availability string   `json:"availability"`
			} `json:"response"`
		}
		var result Result
		body := callAPIWithParams(
			http.MethodGet,
			"/address/"+args[0]+"/availability",
			nil,
			false,
		)
		err := json.Unmarshal(body, &result)
		cobra.CheckErr(err)
		if result.Request.Success {
			fmt.Println(result.Response.Message)
			if result.Response.SeeAlso != nil {
				fmt.Println("See also:")
				for _, seeAlso := range result.Response.SeeAlso {
					fmt.Println("  " + seeAlso)
				}
			}
		} else {
			cobra.CheckErr(fmt.Errorf(result.Response.Message))
		}
	},
}

func init() {
	getAddressCmd.AddCommand(getAddressAvailabilityCmd)
}
