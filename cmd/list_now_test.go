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

	"golang.org/x/exp/slices"
)

func Test_listNow(t *testing.T) {
	result, err := listNow()
	if err != nil {
		t.Errorf("listNow() error = %v", err)
		return
	}
	var returnedAddresses []string
	for _, address := range result.Response.Garden {
		returnedAddresses = append(returnedAddresses, address.Address)
	}
	// NOTE: assumes no listed Now page
	if slices.Contains(returnedAddresses, os.Getenv("CLILOL_ADDRESS")) {
		t.Errorf("listNow() = %v, want %v", returnedAddresses, os.Getenv("CLILOL_ADDRESS"))
	}
}
