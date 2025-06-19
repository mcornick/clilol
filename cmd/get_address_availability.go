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

type getAddressAvailabilityOutput struct {
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

var getAddressAvailabilityCmd = &cobra.Command{
	Use:   "availability [address]",
	Short: "Get address availability",
	Long:  "Gets the availability of an address.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := getAddressAvailability(args[0])
		if err != nil {
			return err
		}
		if result.Request.Success {
			fmt.Println(result.Response.Message)
			if result.Response.SeeAlso != nil {
				fmt.Println("See also:")
				for _, seeAlso := range result.Response.SeeAlso {
					fmt.Println("  " + seeAlso)
				}
			}
		} else {
			return fmt.Errorf(result.Response.Message)
		}
		return nil
	},
}

func init() {
	getAddressCmd.AddCommand(getAddressAvailabilityCmd)
}

func getAddressAvailability(address string) (getAddressAvailabilityOutput, error) {
	var result getAddressAvailabilityOutput
	body := callAPIWithParams(
		http.MethodGet,
		"/address/"+address+"/availability",
		nil,
		false,
	)
	err := json.Unmarshal(body, &result)
	return result, err
}
