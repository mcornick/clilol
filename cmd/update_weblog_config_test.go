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

func Test_updateWeblogConfig(t *testing.T) {
	updateResult, err := updateWeblogConfig("testdata/weblog_config.txt")
	if err != nil {
		t.Errorf("updateWeblogConfig() error = %v", err)
		return
	}
	expected := "Your weblog configuration has been updated. Your weblog is rebuilding, which may take some extra time to complete."
	if updateResult.Response.Message != expected {
		t.Errorf("updateWeblogConfig() = %v, want %v", updateResult.Response.Message, expected)
	}
}
