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

func Test_updateWeblogTemplate(t *testing.T) {
	t.Parallel()
	updateResult, err := updateWeblogTemplate("testdata/weblog_template.txt")
	if err != nil {
		t.Errorf("updateWeblogTemplate() error = %v", err)
		return
	}
	expected := "Your weblog template has been updated."
	if updateResult.Response.Message != expected {
		t.Errorf("updateWeblogTemplate() = %v, want %v", updateResult.Response.Message, expected)
	}
}
