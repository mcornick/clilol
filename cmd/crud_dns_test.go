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

func Test_crudDNS(t *testing.T) {
	expectedName := "localhost." + os.Getenv("CLILOL_ADDRESS")
	expectedType := "A"
	expectedData := "127.0.0.1"
	err := createDNS("localhost", expectedType, expectedData, 0, 3600)
	if err != nil {
		t.Errorf("createDNS() error = %v", err)
		return
	}

	_, err = listDNS()
	if err != nil {
		t.Errorf("listDNS() error = %v", err)
		return
	}

	getResult, err := getDNS(expectedName, expectedType, expectedData, 0, 3600)
	if err != nil {
		t.Errorf("getDNS() error = %v", err)
		return
	}
	recordID := getResult.ID

	err = updateDNS(recordID, "localghost", expectedType, expectedData, 0, 3600)
	if err != nil {
		t.Errorf("updateDNS() error = %v", err)
		return
	}

	err = deleteDNS(recordID)
	if err != nil {
		t.Errorf("deleteDNS() error = %v", err)
		return
	}
}
