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
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type updateNowInput struct {
	Content string `json:"content"`
	Listed  int    `json:"listed"`
}

type updateNowOutput struct {
	Response struct {
		Message string `json:"message"`
	} `json:"response"`
	Request resultRequest `json:"request"`
}

var (
	updateNowFilename string
	updateNowListed   bool
	updateNowCmd      = &cobra.Command{
		Use:   "now",
		Short: "update Now page content",
		Long: `Sets your Now page content.

If you specify a filename with the --filename flag, the content of the file
will be used. If you do not specify a filename, the content will be read
from stdin.

The Now page will be created as unlisted by default. To create a listed
Now page, use the --listed flag.`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			result, err := updateNow(updateNowFilename, updateNowListed)
			cobra.CheckErr(err)
			if result.Request.Success {
				fmt.Println(result.Response.Message)
			} else {
				cobra.CheckErr(fmt.Errorf("%s", result.Response.Message))
			}
		},
	}
)

func init() {
	updateNowCmd.Flags().StringVarP(
		&updateNowFilename,
		"filename",
		"f",
		"",
		"file to read Now page from (default stdin)",
	)
	updateNowCmd.Flags().BoolVarP(
		&updateNowListed,
		"listed",
		"l",
		false,
		"create Now page as listed (default false)",
	)
	updateCmd.AddCommand(updateNowCmd)
}

func updateNow(filename string, listed bool) (updateNowOutput, error) {
	var result updateNowOutput
	var listedB int
	var content string
	if filename != "" {
		updateNowInput, err := os.ReadFile(filename)
		cobra.CheckErr(err)
		content = string(updateNowInput)
	} else {
		stdin, err := io.ReadAll(os.Stdin)
		cobra.CheckErr(err)
		content = string(stdin)
	}
	if listed {
		listedB = 1
	} else {
		listedB = 0
	}
	nowPage := updateNowInput{content, listedB}
	body := callAPIWithParams(
		http.MethodPost,
		"/address/"+viper.GetString("address")+"/now",
		nowPage,
		true,
	)
	err := json.Unmarshal(body, &result)
	return result, err
}
