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

func Test_getWeblogTemplate(t *testing.T) {
	result, err := getWeblogTemplate()
	if err != nil {
		t.Errorf("getWeblogTemplate() error = %v", err)
		return
	}
	expected := resultRequest{StatusCode: 200, Success: true}
	if !reflect.DeepEqual(result.Request, expected) {
		t.Errorf("getWeblogTemplate() = %v, want %v", result.Request, expected)
	}
}
