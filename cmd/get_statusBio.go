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

var getStatusBioCmd = &cobra.Command{
	Use:   "status-bio",
	Short: "Get status bio",
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
		if addressFlag == "" {
			addressFlag = viper.GetString("address")
		}
		body := callAPIWithParams(
			http.MethodGet,
			"/address/"+addressFlag+"/statuses/bio/",
			nil,
			false,
		)
		err := json.Unmarshal(body, &result)
		checkError(err)
		if !silentFlag {
			if !jsonFlag {
				if result.Request.Success {
					fmt.Println(result.Response.Message)
					fmt.Printf("\n%s\n", result.Response.Bio)
				} else {
					checkError(fmt.Errorf(result.Response.Message))
				}
			} else {
				fmt.Println(string(body))
			}
		}
	},
}

func init() {
	getStatusBioCmd.Flags().StringVarP(
		&addressFlag,
		"address",
		"a",
		"",
		"address whose status bio to get",
	)
	getCmd.AddCommand(getStatusBioCmd)
}
