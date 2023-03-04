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

func Test_updateWeb(t *testing.T) {
	content, err := os.ReadFile("testdata/web.txt")
	if err != nil {
		t.Errorf("os.ReadFile() error = %v", err)
		return
	}
	publish := true
	published, err := updateWeb(content, publish)
	if err != nil {
		t.Errorf("updateWeb() error = %v", err)
		return
	}
	if !published {
		t.Errorf("updateWeb() published = %v, want %v", published, true)
	}
}
