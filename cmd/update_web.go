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

type updateWebInput struct {
	Publish bool   `json:"publish,omitempty"`
	Content string `json:"content"`
}

type updateWebOutput struct {
	Request  resultRequest `json:"request"`
	Response struct {
		Message string `json:"message"`
	} `json:"response"`
}

var (
	updateWebFilename string
	updateWebPublish  bool
	updateWebCmd      = &cobra.Command{
		Use:   "web",
		Short: "set webpage content",
		Long: `Sets your webpage content.

If you specify a filename with the --filename flag, the content of the file
will be used. If you do not specify a filename, the content will be read
from stdin.

The webpage will be created as unpublished by default. To create a published
webpage, use the --publish flag.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := updateWeb(updateWebFilename)
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
	updateWebCmd.Flags().StringVarP(
		&updateWebFilename,
		"filename",
		"f",
		"",
		"file to read webpage from (default stdin)",
	)
	updateWebCmd.Flags().BoolVarP(
		&updateWebPublish,
		"publish",
		"p",
		false,
		"publish the updated page (default false)",
	)

	updateCmd.AddCommand(updateWebCmd)
}

func updateWeb(filename string) (updateWebOutput, error) {
	var result updateWebOutput
	var content string
	if filename != "" {
		updateWebInput, err := os.ReadFile(filename) // #nosec G304
		if err != nil {
			return result, err
		}
		content = string(updateWebInput)
	} else {
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			return result, err
		}
		content = string(stdin)
	}
	webPage := updateWebInput{updateWebPublish, content}
	body, err := callAPIWithParams(
		http.MethodPost,
		"/address/"+viper.GetString("address")+"/web",
		webPage,
		true,
	)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}
