// SPDX-License-Identifier: MPL-2.0
//
// Copyright (c) 2023 Mark Cornick
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type updateWebPFPOutput struct {
	Response struct {
		Message string `json:"message"`
	} `json:"response"`
	Request resultRequest `json:"request"`
}

var updateWebPFPCmd = &cobra.Command{
	Use:   "pfp [filename]",
	Short: "set your profile picture",
	Long:  "Sets your profile picture.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := updateWebPFP(args[0])
		if err != nil {
			return err
		}
		if result.Request.Success {
			fmt.Println(result.Response.Message)
		} else {
			return fmt.Errorf("%s", result.Response.Message)
		}
		return nil
	},
}

func init() {
	updateWebCmd.AddCommand(updateWebPFPCmd)
}

func updateWebPFP(filename string) (updateWebPFPOutput, error) {
	var result updateWebPFPOutput
	content, err := os.ReadFile(filename)
	if err != nil {
		return result, err
	}
	encoded := "data:text/plain;base64," + base64.StdEncoding.EncodeToString(content)
	body, err := callAPIWithRawData(
		http.MethodPost,
		"/address/"+viper.GetString("address")+"/pfp",
		encoded,
		true,
	)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}
