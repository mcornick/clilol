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

func Test_updateWebPFP(t *testing.T) {
	t.Parallel()
	updateResult, err := updateWebPFP("testdata/pfp.gif")
	if err != nil {
		t.Errorf("updateWebPFP() error = %v", err)
		return
	}
	expected := "Your image was saved."
	if updateResult.Response.Message != expected {
		t.Errorf("updateWebPFP() = %v, want %v", updateResult.Response.Message, expected)
	}
}
