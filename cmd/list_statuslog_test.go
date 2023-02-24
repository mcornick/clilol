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

func Test_listStatuslog(t *testing.T) {
	result, err := listStatuslog(false)
	if err != nil {
		t.Errorf("listStatuslog() error = %v", err)
		return
	}
	var returnedAddresses []string
	for _, address := range result.Response.Statuses {
		returnedAddresses = append(returnedAddresses, address.Address)
	}
	if !slices.Contains(returnedAddresses, os.Getenv("CLILOL_ADDRESS")) {
		t.Errorf("listStatuslog() = %v, want %v", returnedAddresses, os.Getenv("CLILOL_ADDRESS"))
	}
}
