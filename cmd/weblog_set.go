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

var weblogSetCmd = &cobra.Command{
	Use:   "set",
	Short: "set your weblog config",
	Long:  `The account get group of commands allows you to set your weblog.lol configuration.`,
}

func init() {
	weblogCmd.AddCommand(weblogSetCmd)
}
