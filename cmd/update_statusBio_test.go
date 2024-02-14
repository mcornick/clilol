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
	"os"
	"testing"
)

func Test_updateStatusBio(t *testing.T) {
	updateResult, err := updateStatusBio("This is a testing account for [clilol](https://clilol.readthedocs.io/)")
	if err != nil {
		t.Errorf("updateStatusBio() error = %v", err)
		return
	}
	expected := fmt.Sprintf(
		"OK, the bio on %s.status.lol has been saved. [View it live.](https://status.lol/%s)",
		os.Getenv("CLILOL_ADDRESS"),
		os.Getenv("CLILOL_ADDRESS"),
	)
	if updateResult.Response.Message != expected {
		t.Errorf("updateStatusBio() = %v, want %v", updateResult.Response.Message, expected)
	}
}
