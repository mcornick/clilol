// SPDX-License-Identifier: MPL-2.0
//
// Copyright (c) 2023 Mark Cornick
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"context"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/mcornick/clilol/cmd"
)

func main() {
	err := fang.Execute(context.Background(), cmd.RootCmd, fang.WithVersion(cmd.Version))
	if err != nil {
		os.Exit(1)
	}
}
