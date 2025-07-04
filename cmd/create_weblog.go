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

type createWeblogOutput struct {
	Response struct {
		Message string `json:"message"`
		Entry   struct {
			Metadata struct {
				Date string `json:"date"`
				Slug string `json:"slug"`
			} `json:"metadata"`
			Location string `json:"location"`
			Title    string `json:"title"`
			Type     string `json:"type"`
			Status   string `json:"status"`
			Body     string `json:"body"`
			Source   string `json:"source"`
			Output   string `json:"createWeblogOutput"`
			Entry    string `json:"entry"`
			Date     int64  `json:"date"`
		} `json:"entry"`
	} `json:"response"`
	Request resultRequest `json:"request"`
}

var (
	createWeblogFilename string
	createWeblogCmd      = &cobra.Command{
		Use:   "weblog",
		Short: "Create a weblog entry",
		Long: `Creates an entry in your weblog.

If you specify a filename with the --filename flag, the content of the file
will be used. If you do not specify a filename, the content will be read
from stdin.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := createWeblog(createWeblogFilename)
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
)

func init() {
	createWeblogCmd.Flags().StringVarP(
		&createWeblogFilename,
		"filename",
		"f",
		"",
		"file to read entry from (default stdin)",
	)
	createCmd.AddCommand(createWeblogCmd)
}

func createWeblog(filename string) (createWeblogOutput, error) {
	var result createWeblogOutput
	var content string
	if filename != "" {
		input, err := os.ReadFile(filename) // #nosec G304
		if err != nil {
			return result, err
		}
		content = string(input)
	} else {
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			return result, err
		}
		content = string(stdin)
	}
	body, err := callAPIWithRawData(
		http.MethodPost,
		"/address/"+viper.GetString("address")+"/weblog/entry",
		content,
		true,
	)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}
