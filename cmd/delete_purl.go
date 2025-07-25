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

type deletePURLOutput struct {
	Request  resultRequest `json:"request"`
	Response struct {
		Message string `json:"message"`
	} `json:"response"`
}

var deletePURLCmd = &cobra.Command{
	Use:   "purl [name]",
	Short: "Delete a PURL",
	Long: `Deletes a PURL.

Note that you won't be asked to confirm deletion.
Be sure you know what you're doing.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := deletePURL(args[0])
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
	deleteCmd.AddCommand(deletePURLCmd)
}

func deletePURL(name string) (deletePURLOutput, error) {
	var result deletePURLOutput
	body, err := callAPIWithParams(
		http.MethodDelete,
		"/address/"+viper.GetString("address")+"/purl/"+name,
		nil,
		true,
	)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}
