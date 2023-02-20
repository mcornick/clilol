// SPDX-License-Identifier: MPL-2.0
//
// Copyright (c) 2023 Mark Cornick
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"github.com/spf13/cobra"
)

var purlCmd = &cobra.Command{
	Use:   "purl",
	Short: "do things with PURLs",
	Long:  `The purl group of commands allows you to get, post, update, and delete PURLs.`,
}

func init() {
	rootCmd.AddCommand(purlCmd)
}
