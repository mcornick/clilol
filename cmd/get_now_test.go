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
	"reflect"
	"testing"
)

func Test_getNow(t *testing.T) {
	result, err := getNow(os.Getenv("CLILOL_ADDRESS"))
	if err != nil {
		t.Errorf("getNow() error = %v", err)
		return
	}
	expected := resultRequest{StatusCode: 200, Success: true}
	if !reflect.DeepEqual(result.Request, expected) {
		t.Errorf("getNow() = %v, want %v", result.Request, expected)
	}
}
