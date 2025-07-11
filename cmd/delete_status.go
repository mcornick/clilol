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

type deleteStatusOutput struct {
	Response struct {
		Message string `json:"message"`
	} `json:"response"`
	Request resultRequest `json:"request"`
}

var deleteStatusCmd = &cobra.Command{
	Use:   "status [id]",
	Short: "Delete a status",
	Long: `Deletes a status.

Note that you won't be asked to confirm deletion.
Be sure you know what you're doing.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := deleteStatus(args[0])
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
	deleteCmd.AddCommand(deleteStatusCmd)
}

func deleteStatus(id string) (deleteStatusOutput, error) {
	var result deleteStatusOutput
	body, err := callAPIWithParams(
		http.MethodDelete,
		"/address/"+viper.GetString("address")+"/statuses/"+id,
		nil,
		true,
	)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}
