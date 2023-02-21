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

var bioGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get status bio",
	Long: `Gets status bio for a user from status.lol.

The address can be specified with the --address flag. If not set,
it defaults to your own address.

Note that any custom CSS set on the bio is ignored.
`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		type Result struct {
			Request struct {
				StatusCode int  `json:"status_code"`
				Success    bool `json:"success"`
			} `json:"request"`
			Response struct {
				Message string `json:"message"`
				Bio     string `json:"bio"`
				Css     string `json:"css"`
			} `json:"response"`
		}
		var result Result
		body := callAPIWithJSON(
			http.MethodGet,
			"/address/"+address+"/statuses/bio/",
			nil,
			false,
		)
		err := json.Unmarshal(body, &result)
		cobra.CheckErr(err)
		if !silent {
			if !wantJson {
				if result.Request.Success {
					fmt.Println(result.Response.Message)
					fmt.Printf("\n%s\n", result.Response.Bio)
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
	bioGetCmd.Flags().StringVarP(
		&address,
		"address",
		"a",
		viper.GetString("address"),
		"address whose status bio to get",
	)
	statusBioCmd.AddCommand(bioGetCmd)
}
