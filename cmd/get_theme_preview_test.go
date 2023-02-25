// SPDX-License-Identifier: MPL-2.0
//
// Copyright (c) 2023 Mark Cornick
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"reflect"
	"testing"
)

func Test_getThemePreview(t *testing.T) {
	t.Parallel()
	result, err := getThemePreview("Default")
	if err != nil {
		t.Errorf("getThemePreview() error = %v", err)
		return
	}
	expected := "Here’s an HTML preview of the Default theme."
	if !reflect.DeepEqual(result.Response.Message, expected) {
		t.Errorf("getThemePreview() = %v, want %v", result.Response.Message, expected)
	}
}
