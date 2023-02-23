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

var getAddressAvailabilityCmd = &cobra.Command{
	Use:   "availability",
	Short: "Get address availability",
	Long:  `Gets the availability of an address.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		type Result struct {
			Request  Request `json:"request"`
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
		if addressFlag == "" {
			addressFlag = viper.GetString("address")
		}
		body := callAPIWithParams(
			http.MethodGet,
			"/address/"+addressFlag+"/availability",
			nil,
			false,
		)
		err := json.Unmarshal(body, &result)
		checkError(err)
		if result.Request.Success {
			log.Info(result.Response.Message)
			if result.Response.SeeAlso != nil {
				fmt.Println("See also:")
				for _, seeAlso := range result.Response.SeeAlso {
					fmt.Println("  " + seeAlso)
				}
			}
		} else {
			checkError(fmt.Errorf(result.Response.Message))
		}
	},
}

func init() {
	getAddressAvailabilityCmd.Flags().StringVarP(
		&addressFlag,
		"address",
		"a",
		"",
		"address whose availability to get",
	)
	getAddressCmd.AddCommand(getAddressAvailabilityCmd)
}
