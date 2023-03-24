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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type getNowOutput struct {
	Response struct {
		Message string `json:"message"`
		Now     struct {
			Content string `json:"content"`
			Updated int64  `json:"updated"`
			Listed  int    `json:"listed"`
		} `json:"now"`
	} `json:"response"`
	Request resultRequest `json:"request"`
}

var (
	getNowFilename string
	getNowCmd      = &cobra.Command{
		Use:   "now",
		Short: "Get a Now page",
		Long: `Gets a Now page by address.

The address can be specified with the --address flag. If not set,
it defaults to your own address.

If you specify a filename with the --filename flag, the content will be written
to that file. If you do not specify a filename, the content will be written
to stdout.`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			validateConfig()
			result, err := getNow(addressFlag)
			cobra.CheckErr(err)
			if result.Request.Success {
				if getNowFilename != "" {
					err = os.WriteFile(getNowFilename, []byte(result.Response.Now.Content), 0o644)
					cobra.CheckErr(err)
				} else {
					fmt.Println(result.Response.Now.Content)
				}
			} else {
				cobra.CheckErr(fmt.Errorf(result.Response.Message))
			}
		},
	}
)

func init() {
	getNowCmd.Flags().StringVarP(
		&addressFlag,
		"address",
		"a",
		"",
		"address whose Now page to get",
	)
	getNowCmd.Flags().StringVarP(
		&getNowFilename,
		"filename",
		"f",
		"",
		"file to write Now page to (default stdout)",
	)
	getCmd.AddCommand(getNowCmd)
}

func getNow(address string) (getNowOutput, error) {
	var result getNowOutput
	if address == "" {
		address = viper.GetString("address")
	}
	body := callAPIWithParams(
		http.MethodGet,
		"/address/"+address+"/now",
		nil,
		false,
	)
	err := json.Unmarshal(body, &result)
	return result, err
}
