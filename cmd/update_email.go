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
	updateEmailDestination string
	updateEmailCmd         = &cobra.Command{
		Use:   "email",
		Short: "set email forwarding address(es)",
		Long: `Sets your email forwarding address(es).
	
To specify multiple addresses, separate them with commas.`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			type Input struct {
				Destination string `json:"destination"`
			}
			type Result struct {
				Request struct {
					StatusCode int  `json:"status_code"`
					Success    bool `json:"success"`
				} `json:"request"`
				Response struct {
					Message           string   `json:"message"`
					DestinationString string   `json:"destination_string"`
					DestinationArray  []string `json:"destination_array"`
					Address           string   `json:"address"`
					EmailAddress      string   `json:"email_address"`
				} `json:"response"`
			}
			var result Result
			email := Input{updateEmailDestination}
			body := callAPIWithJSON(
				http.MethodPost,
				"/address/"+viper.GetString("address")+"/email",
				email,
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
	updateEmailCmd.Flags().StringVarP(
		&updateEmailDestination,
		"destination",
		"d",
		"",
		"address(es) to forward to",
	)
	updateCmd.AddCommand(updateEmailCmd)
}
