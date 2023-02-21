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

var accountSetCmd = &cobra.Command{
	Use:   "set",
	Short: "set information about your account",
	Long:  `The account set group of commands allows you to set information about your omg.lol account.`,
}

func init() {
	accountCmd.AddCommand(accountSetCmd)
}
