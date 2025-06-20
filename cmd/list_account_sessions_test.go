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

func Test_listAccountSessions(t *testing.T) {
	result, err := listAccountSessions()
	if err != nil {
		t.Errorf("listAccountSessions() error = %v", err)
		return
	}
	if !result.Request.Success {
		t.Errorf("listAccountSessions() result.Request.Success = %v, want %v", result.Request.Success, true)
	}
}
