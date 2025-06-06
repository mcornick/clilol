// SPDX-License-Identifier: MPL-2.0
//
// Copyright (c) 2023 Mark Cornick
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"os"

	"github.com/mcornick/omglol-client-go/omglol"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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
		Run: func(cmd *cobra.Command, args []string) {
			content, err := getWeb()
			handleAPIError(err)
			var writeErr error
			if getWebFilename != "" {
				writeErr = os.WriteFile(getWebFilename, content, 0o644)
			} else {
				_, writeErr = os.Stdout.Write(content)
			}
			cobra.CheckErr(writeErr)
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

func getWeb() ([]byte, error) {
	client, err := omglol.NewClient(viper.GetString("email"), viper.GetString("apikey"), endpoint)
	cobra.CheckErr(err)
	content, err := client.GetWeb(viper.GetString("address"))
	if err != nil {
		return nil, err
	} else {
		return content.ContentBytes, nil
	}
}
