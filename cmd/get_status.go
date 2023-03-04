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
	Request  resultRequest `json:"request"`
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
}

var getStatusCmd = &cobra.Command{
	Use:   "status [id]",
	Short: "Get status",
	Long: `Gets a single status for a single user from status.lol.

The address can be specified with the --address flag. If not set,
it defaults to your own address.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		validateConfig()
		result, err := getStatus(addressFlag, args[0])
		cobra.CheckErr(err)
		if result.Request.Success {
			fmt.Printf(
				"\nhttps://status.lol/%s/%s\n",
				result.Response.Status.Address,
				result.Response.Status.Id,
			)
			timestamp, err := strconv.Atoi(result.Response.Status.Created)
			cobra.CheckErr(err)
			fmt.Printf("  %s\n", time.Unix(int64(timestamp), 0))
			fmt.Printf(
				"  %s %s\n",
				result.Response.Status.Emoji,
				result.Response.Status.Content,
			)
		} else {
			cobra.CheckErr(fmt.Errorf(result.Response.Message))
		}
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
	body := callAPIWithParams(
		http.MethodGet,
		"/address/"+address+"/statuses/"+id,
		nil,
		false,
	)
	err := json.Unmarshal(body, &result)
	return result, err
}
