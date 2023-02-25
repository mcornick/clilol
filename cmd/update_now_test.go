// SPDX-License-Identifier: MPL-2.0
//
// Copyright (c) 2023 Mark Cornick
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"testing"
)

func Test_updateNow(t *testing.T) {
	updateResult, err := updateNow("testdata/now_page.txt", false)
	if err != nil {
		t.Errorf("updateNow() error = %v", err)
		return
	}
	expected := "OK, your /now page has been updated."
	if updateResult.Response.Message != expected {
		t.Errorf("updateNow() = %v, want %v", updateResult.Response.Message, expected)
	}
}
