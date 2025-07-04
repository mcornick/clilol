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
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type getStatusOutput struct {
	Response struct {
		Message string `json:"message"`
		Status  struct {
			Id      string `json:"id"`
			Address string `json:"address"`
			Created string `json:"created"`
			Emoji   string `json:"emoji"`
			Content string `json:"content"`
		} `json:"status"`
	} `json:"response"`
	Request resultRequest `json:"request"`
}

var getStatusCmd = &cobra.Command{
	Use:   "status [id]",
	Short: "Get status",
	Long: `Gets a single status for a single user from status.lol.

The address can be specified with the --address flag. If not set,
it defaults to your own address.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := getStatus(addressFlag, args[0])
		if err != nil {
			return err
		}
		if result.Request.Success {
			fmt.Printf(
				"\nhttps://status.lol/%s/%s\n",
				result.Response.Status.Address,
				result.Response.Status.Id,
			)
			timestamp, err := strconv.Atoi(result.Response.Status.Created)
			if err != nil {
				return err
			}
			fmt.Printf("  %s\n", time.Unix(int64(timestamp), 0))
			fmt.Printf(
				"  %s %s\n",
				result.Response.Status.Emoji,
				result.Response.Status.Content,
			)
		} else {
			return fmt.Errorf("%s", result.Response.Message)
		}
		return nil
	},
}

func init() {
	getStatusCmd.Flags().StringVarP(
		&addressFlag,
		"address",
		"a",
		"",
		"address whose status to get",
	)
	getCmd.AddCommand(getStatusCmd)
}

func getStatus(address string, id string) (getStatusOutput, error) {
	var result getStatusOutput
	if address == "" {
		address = viper.GetString("address")
	}
	body, err := callAPIWithParams(
		http.MethodGet,
		"/address/"+address+"/statuses/"+id,
		nil,
		false,
	)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}
