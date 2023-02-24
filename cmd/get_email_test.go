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

func Test_getEmail(t *testing.T) {
	getResult, err := getEmail()
	if err != nil {
		t.Errorf("getEmail() error = %v", err)
		return
	}
	if getResult.Response.DestinationString != os.Getenv("CLILOL_EMAIL") {
		t.Errorf("getEmail() = %v, want %v", getResult.Response.DestinationString, os.Getenv("CLILOL_EMAIL"))
	}
}
