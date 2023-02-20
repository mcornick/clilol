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

var statuslogCmd = &cobra.Command{
	Use:   "statuslog",
	Short: "view the status log",
	Long:  `The statuslog group of commands allows you to view the status log.`,
}

func init() {
	rootCmd.AddCommand(statuslogCmd)
}
