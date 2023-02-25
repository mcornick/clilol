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
	t.Parallel()
	expectedName := "createdpurl"
	expectedURL := "https://example.com"
	createResult, err := createPURL(expectedName, expectedURL, false)
	if err != nil {
		t.Errorf("createPURL() error = %v", err)
		return
	}
	if createResult.Response.Name != expectedName {
		t.Errorf("createPURL() = %v, want %v", createResult.Response.Name, expectedName)
	}
	if createResult.Response.URL != expectedURL {
		t.Errorf("createPURL() = %v, want %v", createResult.Response.URL, expectedURL)
	}

	listResult, err := listPURL(os.Getenv("CLILOL_ADDRESS"))
	if err != nil {
		t.Errorf("listPURL() error = %v", err)
		return
	}
	var expectedNames []string
	for _, status := range listResult.Response.PURLs {
		expectedNames = append(expectedNames, status.Name)
	}
	if !slices.Contains(expectedNames, expectedName) {
		t.Errorf("listPURL() = %v, want %v", expectedNames, expectedName)
	}

	getResult, err := getPURL(os.Getenv("CLILOL_ADDRESS"), expectedName)
	if err != nil {
		t.Errorf("getPURL() error = %v", err)
		return
	}
	if getResult.Response.PURL.Name != expectedName {
		t.Errorf("getPURL() = %v, want %v", getResult.Response.PURL.Name, expectedName)
	}

	deleteResult, err := deletePURL(expectedName)
	if err != nil {
		t.Errorf("deletePURL() error = %v", err)
		return
	}
	expectedMessage := "OK, that PURL has been deleted."
	if deleteResult.Response.Message != expectedMessage {
		t.Errorf("deletePURL() = %v , want %v", deleteResult.Response.Message, expectedMessage)
	}
}
