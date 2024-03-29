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

func Test_crudStatus(t *testing.T) {
	expectedText := "this is a created status"
	createResult, err := createStatus(expectedText, "🧪", true)
	if err != nil {
		t.Errorf("createStatus() error = %v", err)
		return
	}
	if createResult.Request.StatusCode != 200 {
		t.Errorf("createStatus() = %v, want %v", createResult.Request.StatusCode, 200)
	}
	statusID := createResult.Response.ID

	listResult, err := listStatus(os.Getenv("CLILOL_ADDRESS"), 0)
	if err != nil {
		t.Errorf("listStatus() error = %v", err)
		return
	}
	listResult2, err := listStatus(os.Getenv("CLILOL_ADDRESS"), len(listResult.Response.Statuses)+1)
	if err != nil {
		t.Errorf("listStatus() error = %v", err)
		return
	}
	if len(listResult2.Response.Statuses) != len(listResult.Response.Statuses) {
		t.Errorf("listStatus() = %v, wanted %v", len(listResult2.Response.Statuses), len(listResult.Response.Statuses))
	}

	getResult, err := getStatus(os.Getenv("CLILOL_ADDRESS"), statusID)
	if err != nil {
		t.Errorf("getStatus() error = %v", err)
		return
	}
	if getResult.Response.Status.Content != expectedText {
		t.Errorf("getStatus() = %v, want %v", getResult.Response.Status.Content, expectedText)
	}

	expectedText = "This status was updated by clilol tests."
	updateResult, err := updateStatus(statusID, expectedText, "🧪")
	if err != nil {
		t.Errorf("updateStatus() error = %v", err)
		return
	}
	if updateResult.Request.StatusCode != 200 {
		t.Errorf("updateStatus() = %v, want %v", updateResult.Request.StatusCode, 200)
	}

	deleteResult, err := deleteStatus(statusID)
	if err != nil {
		t.Errorf("deleteStatus() error = %v", err)
		return
	}
	expectedMessage := "OK, that status has been deleted."
	if deleteResult.Response.Message != expectedMessage {
		t.Errorf("deleteStatus() = %v , want %v", deleteResult.Response.Message, expectedMessage)
	}
}
