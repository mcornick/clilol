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

func Test_updateAccountSettings(t *testing.T) {
	t.Parallel()
	updateResult, err := updateAccountSettings("email_ok", "iso_8601", "advanced")
	if err != nil {
		t.Errorf("updateAccountSettings() error = %v", err)
		return
	}
	expected := "OK, your settings have been updated."
	if updateResult.Response.Message != expected {
		t.Errorf("updateAccountSettings() = %v, want %v", updateResult.Response.Message, expected)
	}
}
