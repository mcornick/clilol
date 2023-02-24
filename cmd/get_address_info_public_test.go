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

func Test_getAddressInfoPublic(t *testing.T) {
	result, err := getAddressInfoPublic(os.Getenv("CLILOL_ADDRESS"))
	if err != nil {
		t.Errorf("getAddressInfoPublic() error = %v", err)
		return
	}
	expected := "This address is not near expiration."
	if result.Response.Expiration.Message != expected {
		t.Errorf("getAddressInfoPublic() = %v, want %v", result.Response.Expiration.Message, expected)
	}
}
