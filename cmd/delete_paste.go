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

type deletePasteOutput struct {
	Request  resultRequest `json:"request"`
	Response struct {
		Message string `json:"message"`
	} `json:"response"`
}

var deletePasteCmd = &cobra.Command{
	Use:   "paste [title]",
	Short: "Delete a paste",
	Long: `Deletes a paste.

Specify the paste title with the --title flag.

Note that you won't be asked to confirm deletion.
Be sure you know what you're doing.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		validateConfig()
		result, err := deletePaste(args[0])
		cobra.CheckErr(err)
		if result.Request.Success {
			fmt.Println(result.Response.Message)
		} else {
			cobra.CheckErr(fmt.Errorf(result.Response.Message))
		}
	},
}

func init() {
	deleteCmd.AddCommand(deletePasteCmd)
}

func deletePaste(title string) (deletePasteOutput, error) {
	var result deletePasteOutput
	body := callAPIWithParams(
		http.MethodDelete,
		"/address/"+viper.GetString("address")+"/pastebin/"+title,
		nil,
		true,
	)
	err := json.Unmarshal(body, &result)
	return result, err
}
