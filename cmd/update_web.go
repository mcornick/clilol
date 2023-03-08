// SPDX-License-Identifier: MPL-2.0
//
// Copyright (c) 2023 Mark Cornick
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/ejstreet/omglol-client-go/omglol"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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
		Run: func(cmd *cobra.Command, args []string) {
			validateConfig()
			var content []byte
			var err error
			if updateWebFilename == "" {
				content, err = io.ReadAll(os.Stdin)
			} else {
				content, err = os.ReadFile(updateWebFilename)
			}
			cobra.CheckErr(err)
			published, err := updateWeb(content, updateWebPublish)
			handleAPIError(err)
			if published {
				fmt.Println("Web content saved and published.")
			} else {
				fmt.Println("Web content saved but not published.")
			}
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

func updateWeb(content []byte, publish bool) (bool, error) {
	client, err := omglol.NewClient(viper.GetString("email"), viper.GetString("apikey"), endpoint)
	cobra.CheckErr(err)
	published, err := client.SetWeb(viper.GetString("address"), content, publish)
	return published, err
}
