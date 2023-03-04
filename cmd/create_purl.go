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

	"github.com/ejstreet/omglol-client-go/omglol"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	createPURLListed bool
	createPURLCmd    = &cobra.Command{
		Use:   "purl [name] [url]",
		Short: "Create a PURL",
		Long: `Creates a PURL.

The PURL will be created as unlisted by default. To create a listed
PURL, use the --listed flag.
`,
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			validateConfig()
			name := args[0]
			url := args[1]
			err := createPURL(name, url, createPURLListed)
			handleAPIError(err)
			fmt.Printf("PURL %s created\n", name)
		},
	}
)

func init() {
	createPURLCmd.Flags().BoolVarP(
		&createPURLListed,
		"listed",
		"l",
		false,
		"create as listed (default false)",
	)
	createCmd.AddCommand(createPURLCmd)
}

func createPURL(name string, url string, listed bool) error {
	client, err := omglol.NewClient(viper.GetString("email"), viper.GetString("apikey"), endpoint)
	cobra.CheckErr(err)
	purl := omglol.NewPersistentURL(name, url, listed)
	err = client.CreatePersistentURL(viper.GetString("address"), *purl)
	return err
}
