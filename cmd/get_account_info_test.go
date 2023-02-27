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

func Test_getAccountInfo(t *testing.T) {
	result, err := getAccountInfo()
	if err != nil {
		t.Errorf("getAccountInfo() error = %v", err)
		return
	}
	if result.Email != os.Getenv("CLILOL_EMAIL") {
		t.Errorf("getAccountInfo() = %v, want %v", result.Email, os.Getenv("CLILOL_EMAIL"))
	}
}
