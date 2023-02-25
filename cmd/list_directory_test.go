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
	"testing"

	"golang.org/x/exp/slices"
)

func Test_listDirectory(t *testing.T) {
	t.Parallel()
	result, err := listDirectory()
	if err != nil {
		t.Errorf("listDirectory() error = %v", err)
		return
	}
	// NOTE: assumes the test address is unlisted
	if slices.Contains(result.Response.Directory, os.Getenv("CLILOL_ADDRESS")) {
		t.Errorf("listDirectory() = %v, want %v", result.Response.Directory, os.Getenv("CLILOL_ADDRESS"))
	}
}
