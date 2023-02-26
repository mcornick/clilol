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

func Test_crudPURL(t *testing.T) {
	expectedName := "createdpurl"
	expectedURL := "https://example.com"
	err := createPURL(expectedName, expectedURL, false)
	if err != nil {
		t.Errorf("createPURL() error = %v", err)
		return
	}

	listResult, err := listPURL(os.Getenv("CLILOL_ADDRESS"))
	if err != nil {
		t.Errorf("listPURL() error = %v", err)
		return
	}
	var expectedNames []string
	for _, status := range listResult {
		expectedNames = append(expectedNames, status.Name)
	}
	if !slices.Contains(expectedNames, expectedName) {
		t.Errorf("listPURL() = %v, want %v", expectedNames, expectedName)
		return
	}

	getResult, err := getPURL(os.Getenv("CLILOL_ADDRESS"), expectedName)
	if err != nil {
		t.Errorf("getPURL() error = %v", err)
		return
	}
	if getResult.Name != expectedName {
		t.Errorf("getPURL() = %v, want %v", getResult.Name, expectedName)
		return
	}

	err = deletePURL(expectedName)
	if err != nil {
		t.Errorf("deletePURL() error = %v", err)
		return
	}
}
