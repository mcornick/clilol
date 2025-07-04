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

type getWebOutput struct {
	Request  resultRequest `json:"request"`
	Response struct {
		Message  string `json:"message"`
		Content  string `json:"content"`
		Type     string `json:"type"`
		Theme    string `json:"theme"`
		CSS      string `json:"css"`
		Head     string `json:"head"`
		Verified int64  `json:"verified"`
		PFP      string `json:"pfp"`
		Metadata string `json:"metadata"`
		Branding string `json:"branding"`
		Modified string `json:"modified"`
	} `json:"response"`
}

var (
	getWebFilename string
	getWebCmd      = &cobra.Command{
		Use:   "web",
		Short: "Get your webpage content",
		Long: `Gets the Markdown content for your webpage.

If you specify a filename with the --filename flag, the content will be written
to that file. If you do not specify a filename, the content will be written
to stdout.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := getWeb()
			if err != nil {
				return err
			}
			if result.Request.Success {
				if getWebFilename != "" {
					err = os.WriteFile(getWebFilename, []byte(result.Response.Content), 0o644)
					if err != nil {
						return err
					}
				} else {
					fmt.Println(result.Response.Content)
				}
			} else {
				return fmt.Errorf("%s", result.Response.Message)
			}
			return nil
		},
	}
)

func init() {
	getWebCmd.Flags().StringVarP(
		&getWebFilename,
		"filename",
		"f",
		"",
		"file to write webpage to (default stdout)",
	)
	getCmd.AddCommand(getWebCmd)
}

func getWeb() (getWebOutput, error) {
	var result getWebOutput
	body, err := callAPIWithParams(
		http.MethodGet,
		"/address/"+viper.GetString("address")+"/web",
		nil,
		true,
	)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}
