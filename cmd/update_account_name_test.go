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
)

func Test_updateAccountName(t *testing.T) {
	t.Parallel()
	updateResult, err := updateAccountName(os.Getenv("CLILOL_NAME"))
	if err != nil {
		t.Errorf("updateAccountName() error = %v", err)
		return
	}
	if updateResult.Response.Name != os.Getenv("CLILOL_NAME") {
		t.Errorf("updateAccountName() = %v, want %v", updateResult.Response.Name, os.Getenv("CLILOL_NAME"))
	}
}
