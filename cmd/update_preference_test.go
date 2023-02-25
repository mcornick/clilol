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

func Test_updatePreference(t *testing.T) {
	result, err := updatePreference("foo", "bar")
	if err != nil {
		t.Errorf("updatePreference() error = %v", err)
		return
	}
	expected := "Your preference has been saved."
	if result.Response.Message != expected {
		t.Errorf("updatePreference() = %v, want %v", result.Response.Message, expected)
	}
}
