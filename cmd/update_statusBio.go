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

var updateStatusBioCmd = &cobra.Command{
	Use:   "status-bio [text]",
	Short: "Update your status bio",
	Long: `Updates your status bio on status.lol.
		
Quote the text if it contains spaces.

Note that the omg.lol API does not permit you to change any custom
CSS. You'll need to do that on the website.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		type input struct {
			Content string `json:"content"`
		}
		type output struct {
			Request  resultRequest `json:"request"`
			Response struct {
				Message string `json:"message"`
				URL     string `json:"url"`
			} `json:"response"`
		}
		var result output
		bio := input{args[0]}
		body := callAPIWithParams(
			http.MethodPost,
			"/address/"+viper.GetString("address")+"/statuses/bio/",
			bio,
			true,
		)
		err := json.Unmarshal(body, &result)
		cobra.CheckErr(err)
		if result.Request.Success {
			fmt.Println(result.Response.Message)
		} else {
			cobra.CheckErr(fmt.Errorf(result.Response.Message))
		}
	},
}

func init() {
	updateCmd.AddCommand(updateStatusBioCmd)
}
