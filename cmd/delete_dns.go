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

type deleteDNSOutput struct {
	Request  resultRequest `json:"request"`
	Response struct {
		Message string `json:"message"`
	} `json:"response"`
}

var deleteDNSCmd = &cobra.Command{
	Use:   "dns [id]",
	Short: "Delete a DNS record",
	Long: `Deletes a DNS record.

Note that you won't be asked to confirm deletion.
Be sure you know what you're doing.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := deleteDNS(args[0])
		if err != nil {
			return err
		}
		if result.Request.Success {
			fmt.Println(result.Response.Message)
		} else {
			return fmt.Errorf(result.Response.Message)
		}
		return nil
	},
}

func init() {
	deleteCmd.AddCommand(deleteDNSCmd)
}

func deleteDNS(id string) (deleteDNSOutput, error) {
	var result deleteDNSOutput
	body, err := callAPIWithParams(
		http.MethodDelete,
		"/address/"+viper.GetString("address")+"/dns/"+id,
		nil,
		true,
	)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}
