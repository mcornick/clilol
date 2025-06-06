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

func Test_crudPaste(t *testing.T) {
	expectedTitle := "createdpaste"
	createResult, err := createPaste(expectedTitle, "testdata/create_paste.txt", false)
	if err != nil {
		t.Errorf("createPaste() error = %v", err)
		return
	}
	if createResult.Response.Title != expectedTitle {
		t.Errorf("createPaste() = %v, want %v", createResult.Response.Title, expectedTitle)
	}

	listResult, err := listPaste(os.Getenv("CLILOL_ADDRESS"))
	if err != nil {
		t.Errorf("listPaste() error = %v", err)
		return
	}
	var expectedTitles []string
	for _, status := range listResult.Response.Pastebin {
		expectedTitles = append(expectedTitles, status.Title)
	}
	if !slices.Contains(expectedTitles, expectedTitle) {
		t.Errorf("listPaste() = %v, want %v", expectedTitles, expectedTitle)
	}

	getResult, err := getPaste(os.Getenv("CLILOL_ADDRESS"), expectedTitle)
	if err != nil {
		t.Errorf("getPaste() error = %v", err)
		return
	}
	if getResult.Response.Paste.Title != expectedTitle {
		t.Errorf("getPaste() = %v, want %v", getResult.Response.Paste.Title, expectedTitle)
	}

	deleteResult, err := deletePaste(expectedTitle)
	if err != nil {
		t.Errorf("deletePaste() error = %v", err)
		return
	}
	expectedMessage := "OK, that paste has been deleted."
	if deleteResult.Response.Message != expectedMessage {
		t.Errorf("deletePaste() = %v , want %v", deleteResult.Response.Message, expectedMessage)
	}
}
